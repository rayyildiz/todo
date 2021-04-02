import {FC} from "react";
import { useColorMode, useColorModeValue, IconButton, ButtonProps } from '@chakra-ui/react';
import { FaMoon, FaSun } from 'react-icons/fa';

interface ColorModeSwitcherProps extends ButtonProps {
}

export const ColorModeSwitcher: FC<ColorModeSwitcherProps> = (props ) => {
  const { toggleColorMode } = useColorMode();
  const text = useColorModeValue('dark', 'light');
  const SwitchIcon = useColorModeValue(FaMoon, FaSun);

  return (
      <IconButton
          size="md"
          fontSize="lg"
          aria-label={`Switch to ${text} mode`}
          variant="ghost"
          color="current"
          marginLeft="2"
          onClick={toggleColorMode}
          icon={<SwitchIcon />}
          {...props}
      />
  )
}
