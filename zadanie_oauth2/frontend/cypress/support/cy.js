Cypress.Commands.add('addFirstProductToCart', () => {
  cy.wait(1000);
  cy.contains('button', 'Dodaj do koszyka', { timeout: 10000 }).should('be.visible').first().click();
  
});

Cypress.Commands.add('addProductToCartByIndex', (index) => {
  cy.wait(1000);
  cy.get('button').eq(index).contains('Dodaj do koszyka').click();
});

Cypress.Commands.add('openCart', () => {

  cy.get('nav').contains('Koszyk').click();
});