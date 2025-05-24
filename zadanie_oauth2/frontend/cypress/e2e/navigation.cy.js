describe('Nawigacja', () => {
  beforeEach(() => {
    cy.visit('/');
  });

  it('Link "Produkty" działa', () => {
    cy.openCart();
    cy.get('nav').contains('Produkty').click();
    cy.contains('Produkty').should('exist');
  });

  it('Link "Koszyk" działa', () => {
    cy.openCart();
    cy.get('nav').contains('Koszyk').click();
    cy.contains('Koszyk').should('exist');
  });

  it('Nawigacja działa z każdej podstrony', () => {
    cy.visit('/cart');
    cy.get('nav').contains('Produkty').click();
    cy.contains('Produkty').should('exist');
    cy.get('nav').contains('Koszyk').click();
    cy.contains('Koszyk').should('exist');
  });
});