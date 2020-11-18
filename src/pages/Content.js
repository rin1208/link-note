import React, { useState, useEffect } from 'react'

import Grid from "@material-ui/core/Grid";
import axios from 'axios'
import styled from "@emotion/styled";
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Link from '@material-ui/core/Link';
import { ReactTinyLink } from 'react-tiny-link'

import Typography from '@material-ui/core/Typography';


const Main = styled.div`
display: "flex";
margin: 20;
text-align: center;
`;

const Sub = styled.div`
  text-align: center;
`;

const Footer = styled.footer`
margin-top: auto;
`;



function Content() {

    const [front, setPost] = useState([]);

    useEffect(() => {
        const headers = {
            'Authorization': localStorage.getItem("jwt") ,
            'Uid': localStorage.getItem("uid") 
        }
        axios
            .get(`${'http://localhost:8080/getcontent'}`, {headers: headers})
            .then(results => {
                setPost(results.data)
                console.log(results.data);
            })
    }, []);

    var list = front.map(function (item) {
        return (
            <Grid item xs={6} md={3}>
                <Card >
                    <CardContent>
                        <Typography  >
                            <p>
                                {item.comment}
                            </p>
                        </Typography>
                        <Typography  >
                            <ReactTinyLink
                                cardSize="small"
                                showGraphic={true}
                                maxLine={2}
                                minLine={1}
                                url={item.url}
                            />
                        </Typography>

                    </CardContent>

                    <Sub>
                        <CardActions>
                            <Button size="big" color="primary">
                                Delete
                        </Button>
                        </CardActions>
                    </Sub>

                </Card>

            </Grid>
        );
    });
    return (
        <div>


            <Main>
                <Grid container spacing={1}>
                    <Grid container item xs={12} spacing={2}>
                        {list}

                    </Grid>
                </Grid>
            </Main>

            <Main>
                <Footer>
                    <p>(c)copy right</p>
                </Footer>
            </Main>

        </div>
    );
}

export default Content;