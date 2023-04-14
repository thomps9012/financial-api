import React from 'react'
import GrantSelect from './grantSelect'

describe('<GrantSelect />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<GrantSelect state="" setState={() => {}} />)
    cy.get("h3").should("contain.text", "Grant")
    cy.get("select").should("have.value", "N/A")
    cy.get("select").select("TANF").should("have.value", "TANF")
  })
})