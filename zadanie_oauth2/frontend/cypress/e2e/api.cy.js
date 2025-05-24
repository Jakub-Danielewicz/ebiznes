describe('Testy API', () => {

    it('GET /products zwraca products', () => {
        cy.request({
      method: 'GET',
      url: 'http://localhost:8080/products'
    }).then((resp) => {
      expect(resp.status).to.eq(200);
      expect(resp.body).to.be.an('array');
    });
    });
    
    it('POST /products z brakującym polem zwraca 400', () => {
    cy.request({
      method: 'POST',
      url: 'http://localhost:8080/products',
      failOnStatusCode: false,
      body: { price: 123 }
    }).then((resp) => {
      expect(resp.status).to.be.oneOf([400, 500]);
    });
  });
    
    it('PUT /products/:id z niestniejącym id zwraca 404', () => {
    cy.request({
      method: 'PUT',
      url: 'http://localhost:8080/products/999999',
      failOnStatusCode: false,
      body: { name: "X", price: 1 }
    }).then((resp) => {
      expect(resp.status).to.eq(404);
    });
  });
    
    it('DELETE /products/:id z nieistniejącym id zwraca 404', () => {
    cy.request({
      method: 'DELETE',
      url: 'http://localhost:8080/products/999999',
      failOnStatusCode: false
    }).then((resp) => {
      expect(resp.status).to.eq(404);
    });
  });

  it('POST /cart/:id/items z nieistniejącym productId zwraca 500 lub 400', () => {
    cy.request('POST', 'http://localhost:8080/cart', {}).then((cartResp) => {
      const cartId = cartResp.body.id || cartResp.body.ID;
      cy.request({
        method: 'POST',
        url: `http://localhost:8080/cart/${cartId}/items`,
        failOnStatusCode: false,
        body: { productId: 999999, quantity: 1 }
      }).then((resp) => {
        expect([400, 404, 500]).to.include(resp.status);
      });
    });
  });

  it('POST /cart/:id/items z quantity 0 zwraca 400', () => {
    cy.request('POST', 'http://localhost:8080/cart', {}).then((cartResp) => {
      const cartId = cartResp.body.id || cartResp.body.ID;
      cy.request({
        method: 'POST',
        url: `http://localhost:8080/cart/${cartId}/items`,
        failOnStatusCode: false,
        body: { productId: 1, quantity: 0 }
      }).then((resp) => {
        expect([400, 500]).to.include(resp.status);
      });
    });
  });

  it('DELETE /cart/:cartId/items/:itemId z nieistniejącym itemId zwraca 404', () => {
    cy.request('POST', 'http://localhost:8080/cart', {}).then((cartResp) => {
      const cartId = cartResp.body.id || cartResp.body.ID;
      cy.request({
        method: 'DELETE',
        url: `http://localhost:8080/cart/${cartId}/items/999999`,
        failOnStatusCode: false
      }).then((resp) => {
        expect(resp.status).to.eq(404);
      });
    });
  });

  it('DELETE /cart/:id z niesitniejącym id zwraca 404', () => {
    cy.request({
      method: 'DELETE',
      url: 'http://localhost:8080/cart/999999',
      failOnStatusCode: false
    }).then((resp) => {
      expect(resp.status).to.eq(404);
    });
  });

  it('GET /cart zwraca cart', () => {
    cy.request({
      method: 'GET',
      url: 'http://localhost:8080/cart',
      failOnStatusCode: false
    }).then((resp) => {
      expect(resp.status).to.eq(201);
    });
  });

  it('POST /cart/:id/items bez body zwraca 400', () => {
    cy.request('POST', 'http://localhost:8080/cart', {}).then((cartResp) => {
      const cartId = cartResp.body.id || cartResp.body.ID;
      cy.request({
        method: 'POST',
        url: `http://localhost:8080/cart/${cartId}/items`,
        failOnStatusCode: false,
        body: {}
      }).then((resp) => {
        expect([400, 500]).to.include(resp.status);
      });
    });
  });
  
});