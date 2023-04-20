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

  it('should log in and log out successfully', () => {
    cy.visit('/auth/signup')

    // Fill out the login form

    cy.get("#username").type('myusername')
    cy.get("#password").type('mypassword')
    cy.get('button[type="submit"]').click()

    // Check if the user is redirected to the home page after login
    cy.url().should('include', '/home')
  })
})

describe('Sign Up', () => {
  it('should require email', () => {
      
    cy.visit('/auth/signup')

    // Fill out the login form
    cy.get("#name").type('myName')
    cy.get("#email").type('myemail')
    cy.get("#password").type('matchingPassword1')
    cy.get("#confirm").type('matchingPassword1')
    cy.get('button[type="submit"]').click()

    cy.url().should('include', '/auth/signup')
  })
})

describe('Sign Up', () => {


  it('should match password', () => {
    cy.visit('/auth/signup')

    // Fill out the login form
    cy.get("#name").type('myName')
    cy.get("#email").type('myemail@email.com')
    cy.get("#password").type('matchingPassword1')
    cy.get("#confirm").type('notMatchingPassword1')
    cy.get('button[type="submit"]').click()

    cy.url().should('include', '/auth/signup')
  })
})


describe('Sign Up', () => {
  it('password contains 8 character', () => {
    cy.visit('/auth/signup')

    // Fill out the login form
    cy.get("#email").type('myName')
    cy.get("#email").type('myemail@email.com')
    cy.get("#password").type('shortP1')
    cy.get("#password").type('shortP1')
    cy.get('button[type="submit"]').click()

    cy.url().should('include', '/auth/signup')
  })
})

describe('Sign Up', () => {

  it('password contains upper', () => {
    cy.visit('/auth/signup')

    // Fill out the login form
    cy.get("#name").type('myName')
    cy.get("#email").type('myemail@email.com')
    cy.get("#password").type('lowercasepassword1')
    cy.get("#confirm").type('lowercasepassword1')
    cy.get('button[type="submit"]').click()

    cy.url().should('include', '/auth/signup')
  })
})

describe('Sign Up', () => {
  it('can sign up and log in', () => {
    cy.visit('/auth/signup')

    // Fill out the login form
    cy.get("#name").type('myName')
    cy.get("#email").type('myemail@email.com')
    cy.get("#password").type('correctPassword1')
    cy.get("#password").type('correctPassword1')
    cy.get('button[type="submit"]').click()

    cy.get("#username").type('myemail@email.com')
    cy.get("#passsword").type('correctPassword1')
    cy.get('button[type="submit"]').click()

    // Check if the user is redirected to the home page after login
    cy.url().should('include', '/home')
  })
})
