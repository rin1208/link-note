import React, { useState, useEffect } from 'react'
import TextField from '@material-ui/core/TextField';
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

const Post_field = styled.div`
    display: flex;
    flex-direction: column;
    margin: 20px 0 0 0
`;
const Footer = styled.footer`
width: 95%;
text-align: center;
padding: 30px 0;
position: absolute;
bottom: 0; 
position: fixed;
`;


function Content() {

    const [front, setPost] = useState([]);
    const [url, setUrl] = useState("");
    const [comment, setComment] = useState("");
    const [isRedirect, setIsRedirect] = useState(false);

    const handleUrl = e => {
        setUrl(e.target.value)
    }

    const handcomment = e => {
        setComment(e.target.value)
    }


    const [formData, setFormData] = useState({
        url: "",
        comment: "",
        user_id: ""
    })


    const fileInput = React.createRef()

    const handleSubmit = async () => {
        const data = { url: url, comment: comment, uid: localStorage.getItem("uid") }
        const headers = {
            'Authorization': localStorage.getItem("jwt") ,
            'Uid': localStorage.getItem("uid") 
        }

        await axios.post(`${'http://localhost:8080/post'}`, data,{headers: headers}).then((results) => {
            console.log(results);
            setUrl("") 
        })
        getContent();
    }

    const getContent = async () => {
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
    }


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
        <Main>
            <Main>
                <Grid container spacing={1}>
                    <Grid container item xs={12} spacing={2}>
                        {list}
                    </Grid>
                </Grid>
            </Main>
            <Main>
                <Footer>
                    <Grid container item xs={12} spacing={2}>
                        <Grid item xs={10} md={10}>
                            <Post_field>
                                <TextField label="url" onChange={handleUrl} value={url} />
                                {/* <TextField label="comment" onChange={handcomment} /> */}
                            </Post_field>
                        </Grid>
                        <Grid item xs={2} md={2}>
                            <Button onClick={handleSubmit} variant="outlined">post</Button>
                        </Grid>
                    </Grid>
                </Footer>
            </Main>
        </Main>
    );
}

export default Content;