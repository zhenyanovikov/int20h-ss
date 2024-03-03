import { Box, IconButton } from "@mui/material";
import useClientStore from "../../../store/clientStore";
import { LOCALE } from "../../../constants/localisation";
import UkraineIcon from "../../../assets/icons/icons8-ukraine.svg";
import GreatBritainIcon from "../../../assets/icons/icons8-great-britain.svg";

function LocaleToggleButton() {
  const locale = useClientStore((state) => state.locale);
  const setLocale = useClientStore((state) => state.setLocale);

  return (
    <IconButton onClick={toggleLocale} color="inherit">
      <Box
        component="img"
        sx={{
          width: "1em",
          height: "1em",
        }}
        alt=""
        src={locale === LOCALE.EN ? GreatBritainIcon : UkraineIcon}
      />
    </IconButton>
  );

  function toggleLocale() {
    setLocale(locale === LOCALE.EN ? LOCALE.UK : LOCALE.EN);
  }
}

export default LocaleToggleButton;
