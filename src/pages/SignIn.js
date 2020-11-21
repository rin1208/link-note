import React, { useEffect, useState } from 'react';
import firebase from '../firebase';
import Button from '@material-ui/core/Button';
import SignInScreen from '../components/SignInScreen'
import styled from "@emotion/styled";
const Main = styled.div`
display: "flex";
margin: 20;
text-align: center;
`;

function SignIn() {
	const [loading, setLoading] = useState(true);
	const [user, setUser] = useState(null);

	useEffect(() => {
		console.log("hogehoge");

		firebase.auth().onAuthStateChanged(user => {
			console.log(user);
			setLoading(false)

			setUser(user)



			if (user) {
				user.getIdToken().then(function (idToken) {

					console.log(idToken);
					localStorage.setItem('jwt', idToken)
				});


				localStorage.setItem('uid', user.uid)
			}
		})
	})

	const logout = () => {
		firebase.auth().signOut();
		localStorage.removeItem('uid')
	}

	if (loading) return <Main>loading</Main>;
	return (
		<Main>
			{user ?
				(<Button onClick={logout} variant="contained" color="primary" disableElevatio>Logout</Button>) :
				(<SignInScreen />)
			}
		</Main>
	)
}

export default SignIn;
