import React from "react";

import styles from "./LogInPage.module.scss"
import MainLogo from "components/MainLogo";
import LogInTile from "./components/LogInTile";

const LogInPage: React.FC = () => {
  return (
    <div className={styles.login_page}>
      <div className={styles.logo_wrapper}>
        <MainLogo isLarge={true} />
      </div>

      <LogInTile headerText="Войти" buttonText="Подтвердить" />
      <LogInTile headerText="Зарегестрироваться" buttonText="Подтвердить" />
    </div>
  );
};

export default LogInPage;
