describe('Płatność', () => {
  beforeEach(() => {
    cy.visit('/');
  });

  it('Przejście do płatności z pustym koszykiem', () => {
    cy.openCart();
    cy.contains('Przejdź do płatności').should('not.exist');
  });

  it('Przejście do płatności z produktami', () => {
    cy.addFirstProductToCart();
    cy.openCart();
    cy.contains('Przejdź do płatności').click();
    cy.contains('Płatność').should('exist');
    cy.get('button').contains('Zamów').should('exist');
  });

  it('Przycisk "Zamów" jest widoczny na stronie płatności', () => {
    cy.addFirstProductToCart();
    cy.openCart();
    cy.contains('Przejdź do płatności').click();
    cy.get('button').contains('Zamów').should('exist');
  });

  it('Po wejściu na /payment widoczny jest nagłówek "Płatność"', () => {
    cy.visit('/payment');
    cy.contains('Płatność').should('exist');
  });

  it('Po złożeniu zamówienia pojawia się alert', () => {
    cy.addFirstProductToCart();
    cy.openCart();
    cy.contains('Przejdź do płatności').click();
    cy.get('button').contains('Zamów').click();
    cy.on('window:alert', (txt) => {
      expect(txt).to.match(/złożone/i);
    });
  });

  it('Po złożeniu zamówienia koszyk jest pusty', () => {
    cy.addFirstProductToCart();
    cy.openCart();
    cy.contains('Przejdź do płatności').click();
    cy.get('button').contains('Zamów').click();
    cy.on('window:alert', () => {});
    cy.openCart();
    cy.contains('Koszyk jest pusty').should('exist');
  });

  it('Złożenie zamówienia czyści koszyk', () => {
    cy.addFirstProductToCart();
    cy.openCart();
    cy.contains('Przejdź do płatności').click();
    cy.get('button').contains('Zamów').click();
    cy.on('window:alert', () => {});
    cy.openCart();
    cy.contains('Koszyk jest pusty').should('exist');
  });
});