import Button from "components/Button";
import React from "react";
import AwaitIcon from "../AwaitIcon";
import DoneIcon from "../DoneIcon";

import styles from "./SetTile.module.scss";

type SetTileProps = {
  setName: string,
  author: string,
  isAwait: boolean
}

const SetTile: React.FC<SetTileProps> = ({setName, author, isAwait}) => {
  return (
    <div className={styles.set_tile}>
      <h1 className={styles.tile_header}>{setName}</h1>

      <div className={styles.info_wrapper}>
        <p className={styles.info}>Автор:
          <span className={styles.info_item}>{author}</span>
        </p>

        <div className={styles.status_wrapper}>
          <p className={styles.info}>Статус: 
            <span className={styles.info_item}>
              {isAwait ? "Ожидание" : "Завершено"}
            </span>
          </p>

          {isAwait ? <AwaitIcon /> : <DoneIcon />}
        </div>
      </div>

      <Button acceptMode={true}>Приступить</Button>
    </div>
  );
};

export default SetTile;
