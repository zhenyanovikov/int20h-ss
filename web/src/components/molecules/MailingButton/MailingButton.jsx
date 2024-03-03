import { Box, Button, IconButton, useMediaQuery } from "@mui/material";
import { Email as EmailIcon } from "@mui/icons-material";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import { ROUTE } from "../../../constants/router";

function MailingButton() {
  const { t } = useTranslation();
  const navigate = useNavigate();
  const isLargerThanSm = useMediaQuery((theme) => theme.breakpoints.up("sm"));

  return (
    <Box>
      {isLargerThanSm ? (
        <Button
          variant="outlined"
          color="inherit"
          startIcon={<EmailIcon />}
          onClick={handleCreateMailing}
        >
          {t("molecules.mailingButton.button.create")}
        </Button>
      ) : (
        <IconButton color="inherit" onClick={handleCreateMailing}>
          <EmailIcon />
        </IconButton>
      )}
    </Box>
  );

  function handleCreateMailing() {
    navigate(ROUTE.EMAIL_TEMPLATES);
  }
}

export default MailingButton;
