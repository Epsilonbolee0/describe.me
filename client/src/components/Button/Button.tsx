import React from 'react';
import styles from "./Button.module.scss"

type ButtonProps = {
	children: React.ReactNode,
	acceptMode: boolean
}

const Button: React.FC<ButtonProps> = ({ children, acceptMode }) => {
	return (
		<>
			{acceptMode ?
				<div className={styles.button_accept}>{children}</div> :
				<div className={styles.button_decline}>{children}</div>
			}
		</>
	);
};

export default Button;
