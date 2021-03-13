import {FC, useState} from "react";
import {gql, useMutation, useQuery} from "@apollo/client";
import {TodoList} from "./__generated__/TodoList";
import {Box, Button, Checkbox, CircularProgress, FormControl, FormControlLabel, FormGroup} from "@material-ui/core";
import {Toggle, ToggleVariables} from "./__generated__/Toggle";
import {NewTodoPage} from "./NewTodo";
import {stat} from "fs";


const TODO_LIST_QUERY = gql`
    query TodoList {
        todos {
            id
            content
            completed
        }
    }
`

const TOGGLE_MUTATION = gql`
    mutation Toggle($id:ID!) {
        toggle(id: $id) {
            id
            content
            completed
        }
    }
`;

type TodoPageProps = {}
export const TodoListPage: FC<TodoPageProps> = (props) => {
  const {data, loading, error, refetch} = useQuery<TodoList>(TODO_LIST_QUERY);
  const [doQuery] = useMutation<Toggle, ToggleVariables>(TOGGLE_MUTATION)
  const [open, setOpen] = useState(false);

  const handleClick = async (id: string) => {
    await doQuery({
      variables: {
        id
      }
    })
  }

  const handleCloseDialog = async (status: boolean) => {
    setOpen(false)
    if (status) {
      await refetch();
    }
  }

  if (loading) {
    return <CircularProgress disableShrink/>
  }

  if (error) {
    return (
        <Box color="text.secondary">
          {error.message}
        </Box>
    )
  }


  return (
      <div>
        <FormControl component="fieldset">
          <h3>TODO List</h3>

          <Button variant="contained" color="primary" onClick={() => setOpen(true)}>
            Add Todo
          </Button>
          <NewTodoPage open={open} onClose={handleCloseDialog}/>

          {data?.todos.map(todo => (
              <FormGroup aria-label="position" row key={todo.id}>
                <FormControlLabel
                    value={todo.completed}
                    control={<Checkbox
                        color="secondary"
                        checked={todo.completed}
                        onClick={e => handleClick(todo.id)}
                    />}
                    label={todo.content}
                    labelPlacement="end"
                />
              </FormGroup>)
          )}
        </FormControl>
      </div>
  )
};
