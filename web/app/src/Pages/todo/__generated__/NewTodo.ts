/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: NewTodo
// ====================================================

export interface NewTodo_new {
  __typename: "Todo";
  id: string;
  content: string;
  completed: boolean;
}

export interface NewTodo {
  new: NewTodo_new | null;
}

export interface NewTodoVariables {
  content: string;
}
