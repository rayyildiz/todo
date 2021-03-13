import React, {FC} from "react";
import {Route, Switch} from "react-router-dom";
import {IndexPage} from "./Pages/IndexPage";
import {PrivacyPage} from "./Pages/PrivacyPage";
import {TodoListPage} from "./Pages/todo/TodoList";

type AppRoutesProps = {}

export const AppRoutes: FC<AppRoutesProps> = (props) => {
  return (
      <Switch>
        <Route exact={true} path="/" component={IndexPage}/>
        <Route path='/privacy' component={PrivacyPage}/>
        <Route path='/todo' component={TodoListPage}/>
      </Switch>
  )
};
