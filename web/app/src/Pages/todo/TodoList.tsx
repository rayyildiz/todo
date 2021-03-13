import {FC, FormEvent} from "react";
import {gql, useMutation, useQuery} from "@apollo/client";
import {TodoList} from "./__generated__/TodoList";
import {Box, Checkbox, CircularProgress, FormControl, FormControlLabel, FormGroup, FormLabel} from "@material-ui/core";
import {Toggle, ToggleVariables} from "./__generated__/Toggle";

type TodoPageProps = {}

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


export const TodoListPage: FC<TodoPageProps> = (props) => {
  const {data, loading, error} = useQuery<TodoList>(TODO_LIST_QUERY);
  const [doQuery] = useMutation<Toggle, ToggleVariables>(TOGGLE_MUTATION)


  const handleClick = async (id: string) => {
    await doQuery({
      variables: {
        id
      }
    })
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
