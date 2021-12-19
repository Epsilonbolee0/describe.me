import React from "react";

import styles from "./ActionTile.module.scss";

import AcceptIcon from "components/AcceptIcon";
import ActionButton from "components/ActionButton";
import ArrowLeftIcon from "components/ArrowLeftIcon";
import ArrowRightIcon from "components/ArrowRightIcon";
import CodeArea from "components/CodeArea";
import DislikeIcon from "components/DislikeIcon";
import DescriptionArea from "components/DescriptionArea";
import LikeIcon from "components/LikeIcon";
import MoveButton from "components/MoveButton";
import { ActionType } from "utils/types";

type ActionTileProps = {
  mode: number,
  lang: string,
  code: string,
  description?: string
};

const ActionTile: React.FC<ActionTileProps> = ({mode, lang, code, description}) => {
  return (
    <div className={styles.action_tile}>
      <div className={styles.action_header}>
        <MoveButton children1={<ArrowLeftIcon />} children2={"Назад"}></MoveButton>

        <h2 className={styles.lang_header}>{`Язык: ${lang}`}</h2>
        
        <MoveButton children1={"Вперёд"} children2={<ArrowRightIcon />}></MoveButton>
      </div>

      <CodeArea language={lang} codeText={code} />

      <div className={styles.action_wrapper}>
        <DescriptionArea mode={mode} description={description}/>

        <div className={styles.button_wrapper}>
          <ActionButton acceptMode={true}>
            {mode === ActionType.eval ? <LikeIcon /> : <AcceptIcon />}
          </ActionButton>

          {mode === ActionType.eval && <ActionButton acceptMode={false}><DislikeIcon /></ActionButton>}
        </div>
      </div>
    </div>
  );
}

export default ActionTile;