import React from "react";

import styles from "./Header.module.scss";
import LogOutIcon from "components/LogOutIcon";
import MainLogo from "components/MainLogo";

type HeaderProps = {
  login: string
}

const Header: React.FC<HeaderProps> = ({login}) => {
  return (
    <div className={styles.header}>
      <MainLogo isLarge={false}/>
      
      <div className={styles.row_wrapper}>
        <button className={styles.login}>{login}</button>
        <button title="Выйти" className={styles.logout_btn}><LogOutIcon /></button>
      </div>
    </div>
  );
};

export default Header;
