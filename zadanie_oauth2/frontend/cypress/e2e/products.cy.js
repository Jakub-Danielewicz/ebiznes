describe('Produkty', () => {
  beforeEach(() => {
    cy.visit('/');
  });

  it('Wyświetla listę produktów', () => {
    cy.contains('Produkty').should('exist');
    cy.get('div').contains('Produkty').should('exist');
    cy.get('button').contains('Dodaj do koszyka').should('exist');
  });

  it('Każdy produkt ma nazwę, cenę i kategorię', () => {
    cy.get('h2').contains('Produkty').parent().within(() => {
      cy.get('div').each(($el) => {
        cy.wrap($el).find('span').should('exist').and('contain.text', 'zł').and('contain.text', 'kategoria:');
      });
    });
  });
 
  });
