import React from 'react';
import SetTile from './components/SetTile';

import styles from "./MainPage.module.scss";

const MainPage: React.FC = () => {
  return (
    <div className={styles.main_page}>
      <SetTile setName='Функции для обработки массивов' author={'HaskellBoi'} isAwait={false} />
      <SetTile setName='Функции для обработки массивов' author={'HaskellBoi'} isAwait={true} />
    </div>
  );
};

export default MainPage;
