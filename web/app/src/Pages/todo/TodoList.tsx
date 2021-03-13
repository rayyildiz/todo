import {FC, FormEvent, useState} from "react";
import {gql, useMutation, useQuery} from "@apollo/client";
import {TodoList} from "./__generated__/TodoList";
import {Backdrop, Checkbox, CircularProgress, Container, createStyles, FormControlLabel, FormGroup, Grid, makeStyles, TextField, Theme} from "@material-ui/core";
import {Toggle, ToggleVariables} from "./__generated__/Toggle";
import {Alert} from "@material-ui/lab";
import {NewTodo, NewTodoVariables} from "./__generated__/NewTodo";


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

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
      paper: {
        marginTop: theme.spacing(),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
      },
      form: {
        width: '100%',
        marginTop: theme.spacing(4),
      },
      submit: {
        margin: theme.spacing(3, 0, 2),
      },
      backdrop: {
        zIndex: theme.zIndex.drawer + 1,
        color: '#fff',
      },
    }),
);

const NEW_TODO_MUTATION = gql`
    mutation NewTodo($content: String!) {
        new(content: $content) {
            id
            content
            completed
        }
    }
`;


type TodoPageProps = {}
export const TodoListPage: FC<TodoPageProps> = (props) => {
  const {data, loading, error, refetch} = useQuery<TodoList>(TODO_LIST_QUERY);
  const [doToggle] = useMutation<Toggle, ToggleVariables>(TOGGLE_MUTATION)
  const [doUpdate] = useMutation<NewTodo, NewTodoVariables>(NEW_TODO_MUTATION);
  const [content, setContent] = useState("");
  const classes = useStyles();

  const handleClick = async (id: string) => {
    await doToggle({
      variables: {
        id
      }
    })
  }

  const handleFormSave = async (e: FormEvent) => {
    e.preventDefault()

    await doUpdate({
      variables: {
        content
      }
    })
    setContent("");
    await refetch();
  }

  return (
      <Container component="main" maxWidth="sm">
        <div className={classes.paper}>
          <form className={classes.form} noValidate={false} onSubmit={handleFormSave}>
            <Grid container spacing={0}>

              {error && <Grid item xs={12}><Alert severity="error">{error.message}</Alert></Grid>}
              <Grid item xs={12}>

                <h3>TODO List</h3>

                <Grid item md={12}>
                  <TextField
                      margin="normal"
                      required
                      fullWidth
                      name="content"
                      label="content"
                      type="text"
                      id="content"
                      autoComplete="off"
                      value={content}
                      autoFocus={true}
                      onChange={(e) => setContent(e.target.value)}
                  />
                </Grid>

                <br/><br/>
                <hr/>

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
              </Grid>
              <Backdrop
                  className={classes.backdrop}
                  open={loading}>
                <CircularProgress color="inherit"/>
              </Backdrop>
            </Grid>
          </form>
        </div>
      </Container>
  )
};
