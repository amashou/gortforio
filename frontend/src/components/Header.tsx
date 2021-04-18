import React, { useState } from 'react'
import { AppBar, IconButton, makeStyles, Toolbar, Typography, Tabs, Tab, Paper } from '@material-ui/core'
import MenuIcon from '@material-ui/icons/Menu'

const useStyle = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    menuButton: {
        marginRight: theme.spacing(2)
    }
}))

const Header = () => {
    const classes = useStyle();
    const [value, setvalue] = useState("coincheck");

    const handleChange = (event: React.ChangeEvent<{}>, newValue: string)  => {
        console.log(newValue)
        console.log(event)
        setvalue(newValue)
    }

    return (
        <div>
            <AppBar>
                <Toolbar>
                    <IconButton edge='start' className={classes.menuButton} color="inherit" aria-label="menu">
                        <MenuIcon />
                    </IconButton>
                    <Typography variant="h6" color="inherit">
                        GortfoRio
                    </Typography>
                </Toolbar>
                <Paper square className={classes.root}>
                    <Tabs indicatorColor="primary"
                        value={value}
                        textColor="primary"
                        aria-label="simple tabs example"
                        onChange={handleChange}
                        centered
                        >
                        <Tab label="Coincheck" value="coincheck" />
                        <Tab label="GMO coin" value="gmocoin" />
                        <Tab label="Zaif" value="zaif" />
                        <Tab label="Binance" value="binance" />
                    </Tabs>
                </Paper>
            </AppBar>
        </div>
    )
}

export default Header
