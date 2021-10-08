import React from "react";

import styles from "./PageSetProcessing.module.scss";
import ActionTile from "components/ActionTile";
import Header from "components/Header";

type PageSetProcessingProps = {

}

const PageSetProcessing: React.FC<PageSetProcessingProps> = () => {
  return (
    <div>
      <Header login="Epsilonbolee0"/>

      <div className={styles.tile_container}>
        <ActionTile evalMode={false}/>
      </div>
    </div>
  );
};

export default PageSetProcessing;
