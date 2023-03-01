describe('Home page test', () => {
  it('checks home page accessible', () => {
    cy.visit('http://localhost:3000/')
  })
})

describe('Button click', () => {
  it('clicks the log in button', () => {
    cy.visit('http://localhost:3000/')

    cy.contains('Log In').click()
  })
})


describe('sign up', () => {
  it('checks sign up accessiblity', () => {
    cy
  })
})