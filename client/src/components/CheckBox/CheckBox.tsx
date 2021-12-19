import React from "react";

import styles from "./CheckBox.module.scss";

type CheckBoxProps = {
  value: string
  labelMargin?: number
}

const CheckBox: React.FC<CheckBoxProps> = ({value, labelMargin = 8}) => {
  return (
    <div className={styles.wrapper}>
      <input
        className={styles.checkbox}
        id={`${value}_checkbox`}
        type="checkbox"
        value={value}
      />

      <label
        className={styles.label}
        htmlFor={`${value}_checkbox`}
        style={{marginLeft: labelMargin}}
      >
        {value}
      </label>
    </div>
  );
};

export default CheckBox;
