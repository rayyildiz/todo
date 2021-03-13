import React from "react";
import {AppBar, Button, createStyles, IconButton, makeStyles, Theme, Toolbar, Typography} from "@material-ui/core";
import HomeIcon from '@material-ui/icons/Home';
import {useHistory} from "react-router-dom";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
      root: {
        flexGrow: 1,
      },
      menuButton: {
        marginRight: theme.spacing(2),
      },
      title: {
        flexGrow: 1,
        display: 'none',
        [theme.breakpoints.up('sm')]: {
          display: 'block',
        },
        cursor: 'pointer',
      },
    }),
);

export const Header = () => {
  const classes = useStyles();
  const history = useHistory();

  return (
      <div className={classes.root}>
        <AppBar position="static">
          <Toolbar>
            <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu" onClick={() => history.push("/")}>
              <HomeIcon/>
            </IconButton>

            <Typography variant="h6" className={classes.title} onClick={() => history.push("/")}>
              TODO App
            </Typography>

            <Button color="inherit" onClick={() => history.push("/todo")}>Todo</Button>
            <Button color="inherit" onClick={() => history.push("/privacy")}>Privacy</Button>
          </Toolbar>
        </AppBar>
      </div>
  )
};
