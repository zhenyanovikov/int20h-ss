import { Tabs, Tab } from "@mui/material";
import { useNavigate } from "react-router-dom";
import { TABS } from "./constants";
import { useTranslation } from "react-i18next";

function GroupTabs() {
  const { t } = useTranslation();
  const navigate = useNavigate();

  return (
    <Tabs value="analytics" variant="scrollable">
      {TABS.map((tab) => (
        <Tab
          key={tab.value}
          value={tab.value}
          label={t(tab.label)}
          onClick={() => handleClick(tab.to)}
        />
      ))}
    </Tabs>
  );

  function handleClick(to) {
    navigate(to);
  }
}

export default GroupTabs;
