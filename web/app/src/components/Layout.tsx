import {Container, Grid} from '@chakra-ui/react';
import {FC} from "react";
import Nav from "./Nav";

type LayoutProps = {
  children: React.ReactNode;
}

export const Layout: FC<LayoutProps> = (props) => {
  return (
      <>
        <Grid gap={6} p={2}>
          <Nav/>
        </Grid>
        <Container maxW="container.xl">
          {props.children}
        </Container>
      </>
  );
}
