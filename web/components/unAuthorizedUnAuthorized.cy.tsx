import React from 'react'
import UnAuthorized from './unAuthorized'

describe('<UnAuthorized />', () => {
  it('renders', () => {
    cy.mount(<UnAuthorized />)
    cy.get('h1').first().should('have.text', 'Unauthorized')
    cy.get('main > :nth-child(3)').should('contain.text', "You are Attempting to Visit")
    cy.get('.archive-btn').should('have.text', 'Back to Safety')
  })
})