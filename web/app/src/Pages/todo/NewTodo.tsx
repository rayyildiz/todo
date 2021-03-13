import {FC, useState} from "react";
import {Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, TextField} from "@material-ui/core";
import {gql, useMutation} from "@apollo/client";
import {NewTodo, NewTodoVariables} from "./__generated__/NewTodo";


const NEW_TODO_MUTATION = gql`
    mutation NewTodo($content: String!) {
        new(content: $content) {
            id
            content
            completed
        }
    }
`;

type TodoPageProps = {
  open: boolean;
  onClose: (saved: boolean) => void;
}
export const NewTodoPage: FC<TodoPageProps> = (props) => {
  const [content, setContent] = useState("");
  const [doQuery, {error}] = useMutation<NewTodo, NewTodoVariables>(NEW_TODO_MUTATION);

  const handleSave = async () => {

    const d = await doQuery({
      variables: {
        content
      }
    })

    console.log(d)
    setContent("");
    props.onClose(true);
  }


  const handleClose = () => {
    props.onClose(false);
  }


  return (
      <Dialog aria-labelledby="simple-dialog-title" open={props.open} onClose={props.onClose}>
        <DialogTitle id="form-dialog-title">Add new Todo</DialogTitle>
        <DialogContent>
          <DialogContentText>
            To add a new todo, enter your content and click save button

            {error && <div>error : {error.message}</div>}
          </DialogContentText>
          <TextField
              autoFocus
              margin="dense"
              id="content"
              label="Your content"
              type="text"
              value={content}
              onChange={(e) => setContent(e.target.value)}
              fullWidth
          />
          <DialogActions>
            <Button onClick={handleClose} color="primary"> Cancel </Button>
            <Button onClick={handleSave} color="primary"> Save </Button>
          </DialogActions>
        </DialogContent>
      </Dialog>
  )
}
