/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: Toggle
// ====================================================

export interface Toggle_toggle {
  __typename: "Todo";
  id: string;
  content: string;
  completed: boolean;
}

export interface Toggle {
  toggle: Toggle_toggle;
}

export interface ToggleVariables {
  id: string;
}
