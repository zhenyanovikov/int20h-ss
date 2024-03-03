import { Button } from "@mui/material";
import { Google as GoogleIcon } from "@mui/icons-material";
import { useTranslation } from "react-i18next";
import { useLogin } from "../../../api/auth";

function GoogleLoginButton() {
  const { t } = useTranslation();

  const handleLogin = useLogin();

  return (
    <Button
      variant="contained"
      onClick={handleLogin}
      startIcon={<GoogleIcon />}
    >
      {t("molecules.googleLoginButton.continueWithGoogle")}
    </Button>
  );
}

export default GoogleLoginButton;
