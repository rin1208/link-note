import React from 'react';
import firebase from '../firebase';
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth';

const uiConfig = {
	signInFlow: 'popup',
	signInSuccessUrl: "/",
	signInOptions: [
		firebase.auth.GoogleAuthProvider.PROVIDER_ID
	],
}

function SignInScreen() {
	return (
		<div>
			<StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebase.auth()} />
		</div>
	);
}

export default SignInScreen;
