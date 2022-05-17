import {FC} from "react";
import {Box, Heading, Button} from "@chakra-ui/react"
import {LOCALSTORAGE_AUTH_KEY} from "../environment";
import {useNavigate} from "react-router-dom";
import {v4 as uuidv4} from "uuid";
import {FaArrowAltCircleRight} from "react-icons/fa";

type IndexPageProps = {}
export const IndexPage: FC<IndexPageProps> = (props) => {
  const navigate = useNavigate();

  const createSession = () => {

    let uuid = localStorage.getItem(LOCALSTORAGE_AUTH_KEY);
    if (uuid === null) {
      uuid = uuidv4();
      localStorage.setItem(LOCALSTORAGE_AUTH_KEY, uuid);
    }
    navigate("/todo");
  }

  return (
        <Box w="100%" p={6}>
          <Heading as="h4" size="md">
            This is a demo application.
            The application opens a temporary session, this to-do lists will be deleted automatically in 2 hours.
            You can access the source code from  <a href="https://go.rayyildiz.dev/todo"><u>this address</u></a>.
          </Heading>
          <Button mt={4} size="lg" onClick={createSession} rightIcon={<FaArrowAltCircleRight />}>Open a session</Button>
        </Box>
  )
};
