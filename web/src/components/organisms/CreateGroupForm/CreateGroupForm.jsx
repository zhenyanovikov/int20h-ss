import {
  Stack,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  TextField,
  Button,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { useFormik } from "formik";
import { useGetFaculties } from "../../../api/faculty";
import { useCreateGroup } from "../../../api/group";
import { getCreateGroupValidationSchema } from "../../../helpers/validation";

function CreateGroupForm() {
  const { t } = useTranslation();
  const { data: facultiesData } = useGetFaculties();
  const { mutate: createGroupMutate, isLoading: isCreateGroupLoading } =
    useCreateGroup();

  const formik = useFormik({
    initialValues: {
      facultyId: "",
      name: "",
    },
    validationSchema: getCreateGroupValidationSchema(),
    onSubmit: handleSubmit,
  });

  const faculties = facultiesData || [];

  return (
    <form onSubmit={formik.handleSubmit}>
      <Stack direction="row" spacing={2}>
        <FormControl size="small">
          <InputLabel id="facultyId">
            {t("organisms.createGroupForm.form.facultyId.label")}
          </InputLabel>
          <Select
            labelId="facultyId"
            id="facultyId"
            value={formik.values.facultyId}
            label={t("organisms.createGroupForm.form.facultyId.label")}
            onChange={formik.handleChange}
          >
            {faculties.map((faculty) => (
              <MenuItem key={faculty.id} value={faculty.id}>
                {faculty.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <TextField
          id="name"
          name="name"
          size="small"
          label={t("organisms.createGroupForm.form.name.label")}
          value={formik.values.name}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.name && !!formik.errors.name}
          helperText={formik.touched.name && t(formik.errors.name)}
        />
        <Stack>
          <Button
            variant="contained"
            type="submit"
            disabled={isCreateGroupLoading}
          >
            {t("organisms.createGroupForm.form.button.submit")}
          </Button>
        </Stack>
      </Stack>
    </form>
  );

  function handleSubmit(values, actions) {
    createGroupMutate(values);

    actions.resetForm();
  }
}

export default CreateGroupForm;
