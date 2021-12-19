import React from "react";
import { ActionType } from "utils/types";

import styles from "./DescriptionArea.module.scss";

type DescriptionAreaProps = {
  mode: number,
  description?: string
};

const MIN_TEXTAREA_HEIGHT = 200;

const DescriptionArea: React.FC<DescriptionAreaProps> = ({
  mode, description=""
}) => {
  const textareaRef = React.useRef<null | HTMLTextAreaElement>(null);
  const [value, setValue] = React.useState("");

  const onChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => setValue(e.target.value);

  React.useLayoutEffect(() => {
    if (textareaRef.current) {
      textareaRef.current.style.height = "inherit";
    
      textareaRef.current.style.height = `${Math.max(
        textareaRef.current?.scrollHeight,
        MIN_TEXTAREA_HEIGHT
      )}px`;
    }
  }, [value]);

  return (
    <textarea 
      onChange={onChange}
      ref={textareaRef}
      placeholder={"Ваше описание..."}
      className={mode ? styles.edit : styles.eval}
      readOnly={!mode}
      defaultValue={mode ? "" : description}
    >
    </textarea>
  );
};

export default DescriptionArea;