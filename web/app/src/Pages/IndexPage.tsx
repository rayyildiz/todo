import {FC} from "react";
import {Button, Container, createStyles, Grid, makeStyles, Theme, Typography} from "@material-ui/core";
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import {LOCALSTORAGE_AUTH_KEY} from "../environment";
import {useHistory} from "react-router-dom";
import {v4 as uuidv4} from "uuid";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
      form: {
        width: '100%',
        marginTop: theme.spacing(12),
      },
    }),
);


type IndexPageProps = {}


export const IndexPage: FC<IndexPageProps> = (props) => {
  const classes = useStyles();
  const history = useHistory();

  const createSession = () => {

    let uuid = localStorage.getItem(LOCALSTORAGE_AUTH_KEY);
    if (uuid === null) {
      uuid = uuidv4();
      localStorage.setItem(LOCALSTORAGE_AUTH_KEY, uuid);
    }
    history.push("/todo");
  }

  return (
      <Container maxWidth="md" className={classes.form}>
        <div>
          <Grid container spacing={0}>
            <Grid item xs={12}>
              <Typography variant="h5" gutterBottom>
                This is a demo application.
                The application opens a temporary session, this session will be deleted automatically after 30 minutes.
                You can access the source code from <a href="https://go.rayyildiz.dev/todo">this address</a>.
              </Typography>
            </Grid>
            <Grid item xs={12}>
              <Button variant="contained" color="primary" size="large" startIcon={<ExitToAppIcon/>} onClick={createSession}>Open a session</Button>
            </Grid>
          </Grid>
        </div>
      </Container>
  )

};
