import React from 'react'
import Loading from './loading'

describe('<Loading />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<Loading />)
    cy.get('main').should('have.class', 'loader-body')
    cy.get('section').should('have.class', 'loader')
  })
})