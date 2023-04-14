import React from 'react'
import Footer from './footer'

describe('<Footer />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<Footer />)
    cy.get("a").should("contain.text", "Designed")
    cy.get("p").first().should("contain.text", new Date().getFullYear())
  })
})