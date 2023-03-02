describe('Home page test', () => {
  it('checks home page accessible', () => {
    cy.visit('/')
  })
})

describe('Login page test', () => {
  it('checks login page accessible', () => {
    cy.visit('/auth/login/')
  })
})

describe('Signup page test', () => {
  it('checks signup page accessible', () => {
    cy.visit('/auth/signup/')
  })
})

describe('Home page login button click', () => {
  it('clicks the log in button', () => {
    cy.visit('/')

    cy.contains('Log In').click()
  })
})

describe('Home page signup button click', () => {
  it('clicks the log in button', () => {
    cy.visit('/')

    cy.contains('Sign Up').click()
  })
})

describe('Login page, signup button click', () => {
  it('clicks the log in button', () => {
    cy.visit('/auth/login/')

    cy.contains('Dont Have an account? Register here.').click()
  })
})

describe('Signup page, already have acct button click', () => {
  it('clicks the log in button', () => {
    cy.visit('/auth/signup/')

    cy.contains('Already have an account? Login here.').click()
  })
})
