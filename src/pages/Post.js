
import React, { useState } from 'react';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { Redirect } from 'react-router-dom'
import axios from 'axios';
import styled from "@emotion/styled";


const Post_field = styled.div`
    display: flex;
    flex-direction: column;
    margin: 20px 0 0 0
`;


function Post() {
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

        await axios.post(`${'http://localhost:8080/post'}`, data).then((results) => {
            console.log(results);
        })
        //これ外せばhomeにリダイレクトする。
        // setIsRedirect(true)
    }



    if (isRedirect === true) {
        return <Redirect to="/" />;
    }

    return (
        <div className="post__wrap">
            <div className="post">

                <Post_field>
                    <TextField label="url" onChange={handleUrl} />
                    <TextField label="comment" onChange={handcomment} />
                </Post_field>
                <div className="post__button">
                    <Button onClick={handleSubmit} variant="outlined">post</Button>
                </div>
            </div>
        </div>
    )
}

export default Post;