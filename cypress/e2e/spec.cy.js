describe('Home page test', () => {
  it('checks home page accessible', () => {
    cy.visit('/')
  })
})

describe('Button click', () => {
  it('clicks the log in button', () => {
    cy.visit('/')

    cy.contains('Log In').click()
  })
})
