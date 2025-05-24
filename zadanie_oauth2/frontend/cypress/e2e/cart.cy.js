before(() => {
  cy.intercept('POST', 'http://localhost:8080/cart', (req) => {
    req.reply((res) => {
      
      res.send({
        statusCode: 200,
        body: {
          cartId: 'generated-cart-id-xyz',
        },
      });


      Cypress.env('cartId', 'generated-cart-id-xyz');
    });
  }).as('createCart');
});



describe('Koszyk', () => {
  beforeEach(() => {
  const cartId = Cypress.env('cartId') || 'fallback-cart-id';
  cy.visit('/', {
    onBeforeLoad(win) {
      win.localStorage.setItem('cartId', cartId);
    },
  });
});

  it('Dodanie produktu do koszyka', () => {
    const cartId = Cypress.env('cartId') || 'fallback-cart-id';

    cy.addFirstProductToCart();
 
    cy.openCart();
    cy.contains('Koszyk').should('exist');
    cy.contains('Usuń').should('exist');
    cy.contains('zł').should('exist');
  });

  it('Usuwanie produktu z koszyka', () => {
    cy.addFirstProductToCart();
    cy.openCart();
    cy.contains('Usuń').should('exist');
    cy.get('button').contains('Usuń').click();
    cy.contains('Koszyk jest pusty').should('exist');
  });

  it('Koszyk nie jest współdzielony między sesjami', () => {
    cy.addFirstProductToCart();
    cy.clearCookies();
    cy.reload();
    cy.openCart();
    cy.contains('Koszyk jest pusty').should('exist');
  });

  it('Można dodać kilka różnych produktów', () => {
    cy.addProductToCartByIndex(0);
    cy.addProductToCartByIndex(1);
    cy.openCart();
    cy.get('button').filter(':contains("Usuń")').should('have.length.at.least', 2);
  });

  it('Można dodać ten sam produkt kilka razy', () => {
    cy.addFirstProductToCart();
    cy.addFirstProductToCart();
    cy.openCart();
    cy.get('div').contains('zł').should('exist');
  });

  it('Koszyk pokazuje poprawną liczbę produktów', () => {
    cy.addProductToCartByIndex(0);
    cy.addProductToCartByIndex(1);
    cy.openCart();
    cy.get('button').filter(':contains("Usuń")').should('have.length.at.least', 2);
  });

  it('Po usunięciu wszystkich produktów koszyk jest pusty', () => {
    cy.addProductToCartByIndex(0);
    cy.addProductToCartByIndex(1);
    cy.openCart();
    cy.wait(1000);
    cy.get('button').filter(':contains("Usuń")').each(($btn) => cy.wrap($btn).click());
    cy.contains('Koszyk jest pusty').should('exist');
  });

  it('Po wejściu na /cart widoczny jest nagłówek "Koszyk"', () => {
    cy.visit('/cart');
    cy.contains('Koszyk').should('exist');
  });
});