import {Link} from 'react-router-dom';
import {Box, Heading, Flex, Text} from '@chakra-ui/react';
import {ColorModeSwitcher} from "./ColorModeSwitcher";


const Nav = () => (
    <Flex
        as="nav"
        align="center"
        justify="space-between"
        wrap="wrap"
        padding="1.5rem"
    >
      <Flex align="center" mr={5}>

        <Link to="/">
            <Heading as="h1" size="lg" letterSpacing={"-.1rem"}>
              Todo App
            </Heading>
        </Link>
      </Flex>

      <Box
          display={{sm: "none", md: "flex"}}
          width={{sm: "full", md: "auto"}}
          alignItems="center"
          flexGrow={1}
      >  </Box>

      <Box
          display={{sm: "none", md: "block"}}
          mt={{base: 4, md: 0}}>
        <Link to="/">
          <Text mt={{base: 4, md: 0}} mr={6} display="block">
            Home
          </Text>
        </Link>
      </Box>
      <Box
          display={{sm: "none", md: "block"}}
          mt={{base: 4, md: 0}}>
        <Link to="/privacy">
          <Text mt={{base: 4, md: 0}} mr={6} display="block">
            Privacy
          </Text>
        </Link>
      </Box>
      <Box
          display={{sm: "none", md: "block"}}
          mt={{base: 4, md: 0}}>


        <ColorModeSwitcher/>
      </Box>
    </Flex>
)

export default Nav;
/*

 */
