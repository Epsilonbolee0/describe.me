import React from 'react';

import styles from "./MoveButton.module.scss";

type MoveButtonProps = {
  children1?: string | React.ReactNode,
  children2: string | React.ReactNode
}

const MoveButton: React.FC<MoveButtonProps> = ({ children1, children2 }) => {

  return (
    <button className={styles.move_button}>
      {children1 && (typeof children1 === "string" ? 
        <p className={`${styles.btn_text} ${styles.left_item}`}>{children1}</p> : 
        <div className={`${styles.btn_icon} ${styles.left_item}`}>{children1}</div>
      )}

      {typeof children2 === "string" ? 
        <p className={styles.btn_text}>{children2}</p> : 
        <div className={styles.btn_icon}>{children2}</div>
      }
    </button>
  );
};

export default MoveButton;