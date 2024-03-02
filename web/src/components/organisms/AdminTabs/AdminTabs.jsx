import { Tabs, Tab } from "@mui/material";
import { useLocation, useNavigate } from "react-router-dom";
import { TABS } from "./constants";
import { useTranslation } from "react-i18next";

function AdminTabs() {
  const { t } = useTranslation();
  const location = useLocation();
  const navigate = useNavigate();

  const selectedTab = TABS.find((tab) => tab.to === location.pathname);

  return (
    <Tabs value={selectedTab?.value} variant="scrollable">
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

export default AdminTabs;
