import {FC, FormEvent, useState} from "react";
import {gql, useMutation, useQuery} from "@apollo/client";
import {TodoList} from "./__generated__/TodoList";
import {Toggle, ToggleVariables} from "./__generated__/Toggle";
import {NewTodo, NewTodoVariables} from "./__generated__/NewTodo";
import {FaTrash} from 'react-icons/fa';
import {FormControl, FormLabel, IconButton, FormHelperText, Divider, Heading, Checkbox, Input, Container, Table, Tbody, Tr, Td} from "@chakra-ui/react"
import {Delete} from "./__generated__/Delete";

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


const NEW_TODO_MUTATION = gql`
    mutation NewTodo($content: String!) {
        new(content: $content) {
            id
            content
            completed
        }
    }
`;

const DELETE_TODO_MUTATION = gql`
    mutation Delete($id:ID!){
        delete(id:$id)
    }
`;

type TodoPageProps = {}
export const TodoListPage: FC<TodoPageProps> = (props) => {
  const {data, refetch} = useQuery<TodoList>(TODO_LIST_QUERY);
  const [doToggle] = useMutation<Toggle, ToggleVariables>(TOGGLE_MUTATION)
  const [doDelete] = useMutation<Delete, ToggleVariables>(DELETE_TODO_MUTATION)
  const [doUpdate] = useMutation<NewTodo, NewTodoVariables>(NEW_TODO_MUTATION);
  const [content, setContent] = useState("");

  const handleToggle = async (id: string) => {
    await doToggle({
      variables: {
        id
      }
    })
  }

  const handleDelete = async (id: string) => {
    await doDelete({
      variables: {
        id
      }
    })
    await refetch();
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
      <Container maxW="container.xl">
        <form onSubmit={handleFormSave}>
          <FormControl id="text" isRequired>
            <FormLabel>Todo:</FormLabel>
            <Input type="text"
                   autoFocus={true}
                   value={content}
                   onChange={(e) => setContent(e.target.value)}
            />
            <FormHelperText>Enter your todo item and press enter to save.</FormHelperText>
          </FormControl>
        </form>
        <Divider mt={8} mb={6}/>

        <Heading mb={4}>Your todo list:</Heading>
        <Table variant="simple">
          <Tbody>
            {data?.todos.map(todo => (
                <Tr key={`tr_${todo.id}`}>
                  <Td>
                    <Checkbox key={`cb_${todo.id}`}
                              size="lg" mt={3}
                              isChecked={todo.completed}
                              onChange={() => handleToggle(todo.id)}>{todo.content}</Checkbox>
                  </Td>
                  <Td>
                    <IconButton
                        size={"xs"}
                        variant="outline"
                        colorScheme="cyan"
                        aria-label="Delete"
                        icon={<FaTrash/>}
                        onClick={() => handleDelete(todo.id)}
                    />
                  </Td>
                </Tr>
            ))}
          </Tbody>
        </Table>
      </Container>
  )
};
