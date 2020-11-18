import React, { useEffect, useState } from 'react';
import firebase from '../firebase';

import SignInScreen from '../components/SignInScreen'

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
				});


				localStorage.setItem('uid', user.uid)
			}
		})
	})

	const logout = () => {
		firebase.auth().signOut();
		localStorage.removeItem('uid')
	}

	if (loading) return <div>loading</div>;
	return (
		<div>
			{user ?
				(<button onClick={logout}>Logout</button>) :
				(<SignInScreen />)
			}
		</div>
	)
}

export default SignIn;
