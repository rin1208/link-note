import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom'
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';


function Header() {
	const uid = localStorage.getItem('uid')
	return (
		<header className="header">

			<Tabs aria-label="simple tabs example">
				<Link to="/">
					<Tab label="Home" />
				</Link>
				{/* <Link to="/post">
					<Tab label="Post" />
				</Link> */}
				<Link to="/login">
					<Tab label="LogOut" />
				</Link>




			</Tabs>
		</header>
	)
}
export default Header;
