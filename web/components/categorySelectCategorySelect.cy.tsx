import React from 'react'
import CategorySelect from './categorySelect'

describe('<CategorySelect />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<CategorySelect state={{
      category: ""
    }} setState={() => {}}/>)
    cy.get("h3").should("contain.text", "Request Category")
    cy.get("select").should("have.value", null)
    cy.get("select").select("ADMINISTRATIVE").should("have.value", "ADMINISTRATIVE")
  })
})