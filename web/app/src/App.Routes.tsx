import React, {FC} from "react";
import {Route, Routes} from "react-router-dom";
import {IndexPage} from "./Pages/IndexPage";
import {PrivacyPage} from "./Pages/PrivacyPage";
import {TodoListPage} from "./Pages/todo/TodoList";

type AppRoutesProps = {}

export const AppRoutes: FC<AppRoutesProps> = (props) => {
  return (
      <Routes>
        <Route path="/" children={<IndexPage />}/>
        <Route path='/privacy' children={<PrivacyPage />}/>
        <Route path='/todo' children={<TodoListPage />}/>
      </Routes>
  )
};
