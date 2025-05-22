package main

import (
	"backend/controllers"
	"backend/db"
	"backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const testDBFile = "shop.db"
const productsEndPoint = "/products"
const cartEndPoint = "/cart"

const errGetProducts = "GetProducts returned error: %v"

func setupTestDB() {
	os.Remove(testDBFile)
	db.ConnectDB()
}

func TestMain(m *testing.M) {
	setupTestDB()
	code := m.Run()
	os.Exit(code)
}

func TestConnectDB(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("ConnectDB paicked: %v", r) //asercja 1
		}
	}()
	db.ConnectDB()
}

func TestGetProductsEmpty(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, productsEndPoint, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := controllers.GetProducts(c)
	if err != nil {
		t.Errorf(errGetProducts, err) //asercja 2
	}
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code) //asercja 3
	}
	var products []models.Product
	err = json.Unmarshal(rec.Body.Bytes(), &products)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
	if len(products) != 0 {
		t.Errorf("Expected 0 products, got %d", len(products)) //asercja 4
	}
}

func TestCreateProducts(t *testing.T) {
	e := echo.New()
	for i := 0; i < 10; i++ {
		product := models.Product{Name: "P" + strconv.Itoa(i), Price: float64(i+1) * 10, CategoryID: 1}
		body, _ := json.Marshal(product)
		req := httptest.NewRequest(http.MethodPost, productsEndPoint, bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if err := controllers.CreateProduct(c); err != nil {
			t.Errorf("CreateProduct returned error: %v", err)
		}
		if rec.Code != http.StatusCreated {
			t.Errorf("CreateProduct %d failed status %d", i, rec.Code) //asercja 5
		}
		var created models.Product
		if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
			t.Errorf("Failed to unmarshal created product: %v", err) //asercja 6
			continue
		}
		assertProductCreated(t, product, created, i)
	}
}

func assertProductCreated(t *testing.T, expected, actual models.Product, i int) {
	if actual.Name != expected.Name {
		t.Errorf("Expected product name %s, got %s", expected.Name, actual.Name) //asercja 7
	}
	if actual.Price != expected.Price {
		t.Errorf("Expected product price %f, got %f", expected.Price, actual.Price) //asercja 8
	}
	if actual.ID == 0 {
		t.Errorf("Expected product ID to be set, got 0") //asercja 9
	}
	if actual.CategoryID != expected.CategoryID {
		t.Errorf("Expected product categoryID %d, got %d", expected.CategoryID, actual.CategoryID) //asercja 10
	}
	if actual.CreatedAt.IsZero() {
		t.Errorf("Expected product CreatedAt to be set, got zero") //asercja 11
	}
	if actual.UpdatedAt.IsZero() {
		t.Errorf("Expected product UpdatedAt to be set, got zero") //asercja 12
	}
	if actual.DeletedAt.Valid {
		t.Errorf("Expected product DeletedAt to be invalid, got valid") //asercja 13
	}
	if actual.CreatedAt != actual.UpdatedAt {
		t.Errorf("Expected product CreatedAt to be equal to UpdatedAt, got different values") //asercja 15
	}
}

func TestGetProducts(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, productsEndPoint, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := controllers.GetProducts(c); err != nil {
		t.Errorf(errGetProducts, err)
	}
	var products []models.Product
	if err := json.Unmarshal(rec.Body.Bytes(), &products); err != nil {
		t.Errorf("Failed to unmarshal products: %v", err)
	}
	if len(products) < 10 {
		t.Errorf("Expected at least 10 products, got %d", len(products)) //asercja 16
	}
	for i, p := range products {
		if p.Name == "" {
			t.Errorf("Product %d has empty name", i) //asercja 17
		}
		if p.Price < 0 {
			t.Errorf("Product %d has negative price", i) //asercja 18
		}
		if p.ID == 0 {
			t.Errorf("Product %d has ID 0", i) //asercja 19
		}
		if p.CategoryID == 0 {
			t.Errorf("Product %d has categoryID 0", i) //asercja 20
		}
		if p.CreatedAt.IsZero() {
			t.Errorf("Product %d has zero CreatedAt timestamp", i) //asercja 21
		}
	}
}

func TestDeleteProduct(t *testing.T) {
	e := echo.New()
	e.DELETE(productsEndPoint+"/:id", controllers.DeleteProduct)
	req := httptest.NewRequest(http.MethodGet, productsEndPoint, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := controllers.GetProducts(c)
	if err != nil {
		t.Errorf(errGetProducts, err)
	}
	var products []models.Product
	err = json.Unmarshal(rec.Body.Bytes(), &products)
	if err != nil {
		t.Errorf("Failed to unmarshal products: %v", err)
	}
	if len(products) == 0 {
		t.Fatalf("No products found to delete") //asercja 23
	}
	productID := products[0].ID
	if strconv.Itoa(int(productID)) == "0" {
		t.Fatalf("Product ID is 0") //asercja 24
	}
	fmt.Println("Deleting product with ID:", strconv.Itoa(int(productID)))
	req = httptest.NewRequest(http.MethodDelete, productsEndPoint+"/"+strconv.Itoa(int(productID)), nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	saveCookies(rec, nil)
	if rec.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", rec.Code) //asercja 25
	}
}

func TestDeleteProductNotFound(t *testing.T) {
	e := echo.New()
	e.DELETE(productsEndPoint+"/:id", controllers.DeleteProduct)
	req := httptest.NewRequest(http.MethodDelete, productsEndPoint+"/99999", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := controllers.DeleteProduct(c)
	if err != nil {
		t.Errorf("DeleteProduct returned error: %v", err) //asercja 24
	}
	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", rec.Code) //asercja 26
	}
}

func TestCartCRUD(t *testing.T) {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.GET(cartEndPoint, controllers.GetOrCreateCart)
	e.POST(productsEndPoint, controllers.CreateProduct)
	e.POST(cartEndPoint+"/:id/items", controllers.AddItemToCart)
	e.GET(cartEndPoint+"/:id", controllers.GetCart)
	e.DELETE(cartEndPoint+"/:cartId/items/:itemId", controllers.RemoveItemFromCart)
	e.DELETE(cartEndPoint+"/:id", controllers.DeleteCart)

	var cookies []*http.Cookie

	cart := createCart(t, e, &cookies)

	productIDs := createProductsForCart(t, e, &cookies, 10)
	if len(productIDs) != 10 {
		t.Fatalf("Expected 10 product IDs, got %d", len(productIDs)) //asercja 27
	}

	addItemsToCart(t, e, &cookies, cart.ID, productIDs)

	cartResp := getCart(t, e, &cookies, cart.ID)
	if len(cartResp.CartItems) != 10 {
		t.Errorf("Expected 10 items in cart, got %d", len(cartResp.CartItems)) //asercja 28
	}

	removeAllItemsFromCart(t, e, &cookies, cart.ID)

	cartResp = getCart(t, e, &cookies, cart.ID)
	if len(cartResp.CartItems) != 0 {
		t.Errorf("Expected 0 items in cart after removal, got %d", len(cartResp.CartItems)) //asercja 29
	}

	deleteCart(t, e, &cookies, cart.ID)
}

func createCart(t *testing.T, e *echo.Echo, cookies *[]*http.Cookie) models.Cart {
	req := httptest.NewRequest(http.MethodGet, cartEndPoint, nil)
	for _, c := range *cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	saveCookies(rec, cookies)

	if rec.Code != http.StatusCreated && rec.Code != http.StatusOK {
		t.Fatalf("Expected status 201 or 200, got %d", rec.Code) //asercja 30
	}
	var cart models.Cart
	if err := json.Unmarshal(rec.Body.Bytes(), &cart); err != nil {
		t.Fatalf("Failed to unmarshal cart: %v", err) //asercja 31
	}
	if cart.ID == 0 {
		t.Fatalf("Cart ID should not be 0") //asercja 32
	}
	if cart.CartItems != nil {
		t.Fatalf("Cart should not have items on creation") //asercja 33
	}
	if cart.UserID != 0 {
		t.Fatalf("Cart should not have user ID on creation") //asercja 34
	}
	if cart.CreatedAt.IsZero() {
		t.Fatalf("Cart should have a valid CreatedAt timestamp") //asercja 35
	}
	if cart.UpdatedAt.IsZero() {
		t.Fatalf("Cart should have a valid UpdatedAt timestamp") //asercja 36
	}
	if cart.DeletedAt.Valid {
		t.Fatalf("Cart should not be deleted on creation") //asercja 37
	}
	return cart
}

func createProductsForCart(t *testing.T, e *echo.Echo, cookies *[]*http.Cookie, n int) []uint {
	productIDs := []uint{}
	for i := 0; i < n; i++ {
		product := models.Product{Name: "CartP" + strconv.Itoa(i), Price: float64(i+1) * 5, CategoryID: 1}
		body, _ := json.Marshal(product)
		req := httptest.NewRequest(http.MethodPost, productsEndPoint, bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		for _, c := range *cookies {
			req.AddCookie(c)
		}
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		saveCookies(rec, cookies)
		if rec.Code != http.StatusCreated {
			t.Errorf("CreateProduct %d failed, got %d", i, rec.Code) //asercja 38
		}
		var created models.Product
		err := json.Unmarshal(rec.Body.Bytes(), &created)
		if err != nil {
			t.Errorf("Failed to unmarshal created product: %v", err)
		}
		productIDs = append(productIDs, created.ID)
	}
	return productIDs
}

func addItemsToCart(t *testing.T, e *echo.Echo, cookies *[]*http.Cookie, cartID uint, productIDs []uint) {
	for i, pid := range productIDs {
		item := map[string]interface{}{"productId": pid, "quantity": i + 1}
		body, _ := json.Marshal(item)
		req := httptest.NewRequest(http.MethodPost, cartEndPoint+"/"+strconv.Itoa(int(cartID))+"/items", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		for _, c := range *cookies {
			req.AddCookie(c)
		}
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		saveCookies(rec, cookies)
		if rec.Code != http.StatusCreated {
			t.Errorf("AddItemToCart %d failed, got %d", i, rec.Code) //asercja 39
		}
		var created models.CartItem
		err := json.Unmarshal(rec.Body.Bytes(), &created)
		if err != nil {
			t.Errorf("Failed to unmarshal created cart item: %v", err)
		}
		if created.ID == 0 {
			t.Errorf("Expected cart item ID to be set, got 0") //asercja 40
		}
		if created.ProductID != pid {
			t.Errorf("Expected cart item product ID %d, got %d", pid, created.ProductID) //asercja 41
		}
		if created.Quantity != i+1 {
			t.Errorf("Expected cart item quantity %d, got %d", i+1, created.Quantity) //asercja 42
		}
		if created.CreatedAt.IsZero() {
			t.Errorf("Expected cart item CreatedAt to be set, got zero") //asercja 43
		}
	}
}

func getCart(t *testing.T, e *echo.Echo, cookies *[]*http.Cookie, cartID uint) models.Cart {
	req := httptest.NewRequest(http.MethodGet, cartEndPoint+"/"+strconv.Itoa(int(cartID)), nil)
	for _, c := range *cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	saveCookies(rec, cookies)
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code) //asercja 46
	}
	var cart models.Cart
	err := json.Unmarshal(rec.Body.Bytes(), &cart)
	if err != nil {
		t.Errorf("Failed to unmarshal cart: %v", err)
	}
	if cart.ID == 0 {
		t.Errorf("Expected cart ID to be set, got 0") //asercja 47
	}
	if cartID != cart.ID {
		t.Errorf("Expected cart ID %d, got %d", cartID, cart.ID) //asercja 48
	}
	if cart.CartItems == nil {
		t.Errorf("Expected cart to have items, got nil") //asercja 49
	}
	if cart.CreatedAt.IsZero() {
		t.Errorf("Expected cart CreatedAt to be set, got zero") //asercja 50
	}
	if cart.UpdatedAt.IsZero() {
		t.Errorf("Expected cart UpdatedAt to be set, got zero") //asercja 51
	}
	if cart.DeletedAt.Valid {
		t.Errorf("Cart should not be deleted, got deleted") //asercja 52
	}

	return cart
}

func removeAllItemsFromCart(t *testing.T, e *echo.Echo, cookies *[]*http.Cookie, cartID uint) {
	cart := getCart(t, e, cookies, cartID)
	if len(cart.CartItems) == 0 {
		t.Errorf("Expected cart to have items") //asercja 53
	}
	for _, item := range cart.CartItems {
		req := httptest.NewRequest(http.MethodDelete, cartEndPoint+"/"+strconv.Itoa(int(cartID))+"/items/"+strconv.Itoa(int(item.ID)), nil)
		for _, c := range *cookies {
			req.AddCookie(c)
		}
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		saveCookies(rec, cookies)
		if rec.Code != http.StatusNoContent {
			t.Errorf("RemoveItemFromCart failed for item %d, got %d", item.ID, rec.Code) //asercja 54
		}
	}
}

func deleteCart(t *testing.T, e *echo.Echo, cookies *[]*http.Cookie, cartID uint) {
	req := httptest.NewRequest(http.MethodDelete, cartEndPoint+"/"+strconv.Itoa(int(cartID)), nil)
	for _, c := range *cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	saveCookies(rec, cookies)
	if rec.Code != http.StatusNoContent {
		t.Errorf("DeleteCart failed, got %d", rec.Code) //asercja 55
	}
}

func saveCookies(rec *httptest.ResponseRecorder, cookies *[]*http.Cookie) {
	for _, setCookie := range rec.Result().Cookies() {
		found := false
		for i, c := range *cookies {
			if c.Name == setCookie.Name {
				(*cookies)[i] = setCookie
				found = true
				break
			}
		}
		if !found {
			*cookies = append(*cookies, setCookie)
		}
	}
}
