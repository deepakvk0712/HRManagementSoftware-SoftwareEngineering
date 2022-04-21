// https://docs.cypress.io/api/table-of-contents

import { mount } from "@vue/test-utils";
// import HRRegisterEmployee from '../../../src/views/HRRegisterEmployee.vue'

describe('Testing login procedure and landing page elements ', () => {
  it('Visits the root of the application at URL ', () => {

    const email = 'dstejas191@hrtools.com'
    const password = '123'

    cy.visit('http://localhost:8081/')
    cy.get('#userNameInput')
      .type('dstejas191@hrtools.com')

    cy.get('#userPasswordInput')
      .type('123')

    cy.contains('Login')
      .wait(10)
      .click()

    cy.request('POST', 'http://localhost:8080/login', {
      email,
      password
    });

    cy.visit('http://localhost:8081/landing')

  });

  // it('HR employee registration form verification ', () => {

  //   const email = 'dstejas191@hrtools.com'
  //   const password = '123'

  //   cy.visit('http://localhost:8081/')
  //   cy.get('#userNameInput')
  //     .type('dstejas191@hrtools.com')

  //   cy.get('#userPasswordInput')
  //     .type('123')

  //   cy.contains('Login')
  //     .wait(10)
  //     .click()

  //   cy.request('POST', 'http://localhost:8080/login', {
  //     email,
  //     password
  //   });

  //   cy.visit('http://localhost:8081/landing')

  //   //Testing the navigation bar here.

  // });

  it('Sends a new notification to a permitted collegue', () => {
    const email = 'dstejas191@hrtools.com'
    const password = '123'

    cy.visit('http://localhost:8081/')
    cy.get('#userNameInput')
      .type('dstejas191@hrtools.com')

    cy.get('#userPasswordInput')
      .type('123')

    cy.contains('Login')
      .wait(10)
      .click()

    cy.request('POST', 'http://localhost:8080/login', {
      email,
      password
    });

    cy.visit('http://localhost:8081/landing')
  });
})
