/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: TodoList
// ====================================================

export interface TodoList_todos {
  __typename: "Todo";
  id: string;
  content: string;
  completed: boolean;
}

export interface TodoList {
  todos: TodoList_todos[];
}
