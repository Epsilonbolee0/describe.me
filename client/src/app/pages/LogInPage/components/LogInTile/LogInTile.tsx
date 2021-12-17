import Button from 'components/Button';
import Input from 'components/Input';
import React from 'react';

import styles from './LogInTile.module.scss';

type LogInTileProps = {
  headerText: string,
  buttonText: string
}

const LogInTile: React.FC<LogInTileProps> = ({headerText, buttonText}) => {
  return (
    <div className={styles.login_tile}>
      <h1 className={styles.tile_header}>{headerText}</h1>
      
      <div className={styles.input_wrapper}>
        <Input placeholder='Логин' type='text' />
        <Input placeholder='Пароль' type='password'/>
      </div>

      <Button acceptMode={true}>{buttonText}</Button>
    </div>
  );
};

export default LogInTile;
