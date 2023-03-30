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

    cy.contains('Login').click()
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

    cy.contains('Need an account?').click()
  })
})

describe('Signup page, already have acct button click', () => {
  it('clicks the log in button', () => {
    cy.visit('/auth/signup/')

    cy.contains('Already have an account?').click()
  })
})
describe('Login and Logout', () => {
  beforeEach(() => {
    cy.visit('/login')
  })

  it('should log in and log out successfully', () => {
    // Fill out the login form
    cy.get('input[name="username"]').type('myusername')
    cy.get('input[name="password"]').type('mypassword')
    cy.get('button[type="submit"]').click()

    // Check if the user is redirected to the home page after login
    cy.url().should('include', '/home')
  })
})
