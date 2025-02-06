module.exports = {
  purge: ['./src/**/*.{html,ts,css,scss}', './projects/*/src/**/*.{html,ts,css,scss}'],
  prefix: '',
  important: ':root',
  separator: ':',
  theme: {
    screens: {
      sm: '640px',
      md: '768px',
      lg: '1024px',
      xl: '1280px',
      xxl: '1600px'
    },
    colors: {
      transparent: 'transparent',

      black: 'var(--text-color)',
      white: 'var(--text-color-light)',
      'contrast-black': 'rgba(0, 0, 0, 0.87)',
      'contrast-white': 'white',

      gray: {
        50: 'var(--color-gray-50)',
        100: 'var(--color-gray-100)',
        200: 'var(--color-gray-200)',
        300: 'var(--color-gray-300)',
        400: 'var(--color-gray-400)',
        500: 'var(--color-gray-500)',
        600: 'var(--color-gray-600)',
        700: 'var(--color-gray-700)',
        800: 'var(--color-gray-800)',
        900: 'var(--color-gray-900)'
      },
      red: {
        50: 'var(--color-red-50)',
        100: 'var(--color-red-100)',
        200: 'var(--color-red-200)',
        300: 'var(--color-red-300)',
        400: 'var(--color-red-400)',
        500: 'var(--color-red-500)',
        600: 'var(--color-red-600)',
        700: 'var(--color-red-700)',
        800: 'var(--color-red-800)',
        900: 'var(--color-red-900)'
      },
      orange: {
        50: 'var(--color-orange-50)',
        100: 'var(--color-orange-100)',
        200: 'var(--color-orange-200)',
        300: 'var(--color-orange-300)',
        400: 'var(--color-orange-400)',
        500: 'var(--color-orange-500)',
        600: 'var(--color-orange-600)',
        700: 'var(--color-orange-700)',
        800: 'var(--color-orange-800)',
        900: 'var(--color-orange-900)'
      },
      'deep-orange': {
        50: 'var(--color-deep-orange-50)',
        100: 'var(--color-deep-orange-100)',
        200: 'var(--color-deep-orange-200)',
        300: 'var(--color-deep-orange-300)',
        400: 'var(--color-deep-orange-400)',
        500: 'var(--color-deep-orange-500)',
        600: 'var(--color-deep-orange-600)',
        700: 'var(--color-deep-orange-700)',
        800: 'var(--color-deep-orange-800)',
        900: 'var(--color-deep-orange-900)'
      },
      amber: {
        50: 'var(--color-amber-50)',
        100: 'var(--color-amber-100)',
        200: 'var(--color-amber-200)',
        300: 'var(--color-amber-300)',
        400: 'var(--color-amber-400)',
        500: 'var(--color-amber-500)',
        600: 'var(--color-amber-600)',
        700: 'var(--color-amber-700)',
        800: 'var(--color-amber-800)',
        900: 'var(--color-amber-900)'
      },
      'light-green': {
        50: 'var(--color-light-green-50)',
        100: 'var(--color-light-green-100)',
        200: 'var(--color-light-green-200)',
        300: 'var(--color-light-green-300)',
        400: 'var(--color-light-green-400)',
        500: 'var(--color-light-green-500)',
        600: 'var(--color-light-green-600)',
        700: 'var(--color-light-green-700)',
        800: 'var(--color-light-green-800)',
        900: 'var(--color-light-green-900)'
      },
      green: {
        50: 'var(--color-green-50)',
        100: 'var(--color-green-100)',
        200: 'var(--color-green-200)',
        300: 'var(--color-green-300)',
        400: 'var(--color-green-400)',
        500: 'var(--color-green-500)',
        600: 'var(--color-green-600)',
        700: 'var(--color-green-700)',
        800: 'var(--color-green-800)',
        900: 'var(--color-green-900)'
      },
      teal: {
        50: 'var(--color-teal-50)',
        100: 'var(--color-teal-100)',
        200: 'var(--color-teal-200)',
        300: 'var(--color-teal-300)',
        400: 'var(--color-teal-400)',
        500: 'var(--color-teal-500)',
        600: 'var(--color-teal-600)',
        700: 'var(--color-teal-700)',
        800: 'var(--color-teal-800)',
        900: 'var(--color-teal-900)'
      },
      cyan: {
        50: 'var(--color-cyan-50)',
        100: 'var(--color-cyan-100)',
        200: 'var(--color-cyan-200)',
        300: 'var(--color-cyan-300)',
        400: 'var(--color-cyan-400)',
        500: 'var(--color-cyan-500)',
        600: 'var(--color-cyan-600)',
        700: 'var(--color-cyan-700)',
        800: 'var(--color-cyan-800)',
        900: 'var(--color-cyan-900)'
      },
      purple: {
        50: 'var(--color-purple-50)',
        100: 'var(--color-purple-100)',
        200: 'var(--color-purple-200)',
        300: 'var(--color-purple-300)',
        400: 'var(--color-purple-400)',
        500: 'var(--color-purple-500)',
        600: 'var(--color-purple-600)',
        700: 'var(--color-purple-700)',
        800: 'var(--color-purple-800)',
        900: 'var(--color-purple-900)'
      },
      'deep-purple': {
        50: 'var(--color-deep-purple-50)',
        100: 'var(--color-deep-purple-100)',
        200: 'var(--color-deep-purple-200)',
        300: 'var(--color-deep-purple-300)',
        400: 'var(--color-deep-purple-400)',
        500: 'var(--color-deep-purple-500)',
        600: 'var(--color-deep-purple-600)',
        700: 'var(--color-deep-purple-700)',
        800: 'var(--color-deep-purple-800)',
        900: 'var(--color-deep-purple-900)'
      },
      primary: {
        50: 'var(--color-primary-50)',
        100: 'var(--color-primary-100)',
        200: 'var(--color-primary-200)',
        300: 'var(--color-primary-300)',
        400: 'var(--color-primary-400)',
        500: 'var(--color-primary-500)',
        600: 'var(--color-primary-600)',
        700: 'var(--color-primary-700)',
        800: 'var(--color-primary-800)',
        900: 'var(--color-primary-900)'
      }
    },
    spacing: {
      px: '1px',
      page: 'var(--padding-page)',
      0: '0',
      1: '0.25rem',
      2: '0.5rem',
      3: '0.75rem',
      4: '1rem',
      5: '1.25rem',
      6: '1.5rem',
      8: '2rem',
      10: '2.5rem',
      12: '3rem',
      14: '3.5rem',
      16: '4rem',
      20: '5rem',
      24: '6rem',
      32: '8rem',
      40: '10rem',
      48: '12rem',
      56: '14rem',
      64: '16rem'
    },
    backgroundColor: theme => ({
      base: 'var(--background-base)',
      card: 'var(--background-card)',
      'app-bar': 'var(--background-app-bar)',
      hover: 'var(--background-hover)',
      ...theme('colors')
    }),
    backgroundPosition: {
      bottom: 'bottom',
      center: 'center',
      left: 'left',
      'left-bottom': 'left bottom',
      'left-top': 'left top',
      right: 'right',
      'right-bottom': 'right bottom',
      'right-top': 'right top',
      top: 'top'
    },
    backgroundSize: {
      auto: 'auto',
      cover: 'cover',
      contain: 'contain'
    },
    borderColor: theme => ({
      ...theme('colors'),
      DEFAULT: 'var(--foreground-divider)',
      color: 'var(--border)'
    }),
    borderRadius: {
      none: '0',
      sm: '0.125rem',
      DEFAULT: '0.25rem',
      lg: '0.5rem',
      full: '9999px'
    },
    borderWidth: {
      DEFAULT: '1px',
      0: '0',
      2: '2px',
      4: '4px',
      8: '8px'
    },
    boxShadow: {
      b: '0 10px 30px 0 rgba(82, 63, 104, .06)',
      DEFAULT: 'var(--elevation-default)',
      1: 'var(--elevation-z1)',
      2: 'var(--elevation-z2)',
      3: 'var(--elevation-z3)',
      4: 'var(--elevation-z4)',
      5: 'var(--elevation-z5)',
      6: 'var(--elevation-z6)',
      7: 'var(--elevation-z7)',
      8: 'var(--elevation-z8)',
      9: 'var(--elevation-z9)',
      10: 'var(--elevation-z10)',
      11: 'var(--elevation-z11)',
      12: 'var(--elevation-z12)',
      13: 'var(--elevation-z13)',
      14: 'var(--elevation-z14)',
      15: 'var(--elevation-z15)',
      16: 'var(--elevation-z16)',
      17: 'var(--elevation-z17)',
      18: 'var(--elevation-z18)',
      19: 'var(--elevation-z19)',
      20: 'var(--elevation-z20)',
      none: 'none'
    },
    container: {
      center: true,
      padding: 'var(--padding-page)'
    },
    cursor: {
      auto: 'auto',
      DEFAULT: 'default',
      pointer: 'pointer',
      wait: 'wait',
      text: 'text',
      move: 'move',
      'not-allowed': 'not-allowed'
    },
    fill: {
      current: 'currentColor'
    },
    flex: {
      1: '1 1 0%',
      auto: '1 1 auto',
      initial: '0 1 auto',
      none: 'none'
    },
    flexGrow: {
      0: '0',
      DEFAULT: '1'
    },
    flexShrink: {
      0: '0',
      DEFAULT: '1'
    },
    fontFamily: {
      sans: ['Open Sans'],
      serif: ['Georgia', 'Cambria', '"Times New Roman"', 'Times', 'serif'],
      mono: ['Menlo', 'Monaco', 'Consolas', '"Liberation Mono"', '"Courier New"', 'monospace']
    },
    fontSize: {
      xs: '0.75rem',
      sm: '0.875rem',
      base: '1rem',
      lg: '1.125rem',
      xl: '1.25rem',
      '2xl': '1.5rem',
      '3xl': '1.875rem',
      '4xl': '2.25rem',
      '5xl': '3rem',
      '6xl': '4rem'
    },
    fontWeight: {
      hairline: '100',
      thin: '200',
      light: '300',
      normal: '400',
      medium: '500',
      semibold: '600',
      bold: '700',
      extrabold: '800',
      black: '900'
    },
    gridTemplateColumns: {
      none: 'none',
      1: 'repeat(1, minmax(0, 1fr))',
      2: 'repeat(2, minmax(0, 1fr))',
      3: 'repeat(3, minmax(0, 1fr))',
      4: 'repeat(4, minmax(0, 1fr))',
      5: 'repeat(5, minmax(0, 1fr))',
      6: 'repeat(6, minmax(0, 1fr))',
      7: 'repeat(7, minmax(0, 1fr))',
      8: 'repeat(8, minmax(0, 1fr))',
      9: 'repeat(9, minmax(0, 1fr))',
      10: 'repeat(10, minmax(0, 1fr))',
      11: 'repeat(11, minmax(0, 1fr))',
      12: 'repeat(12, minmax(0, 1fr))'
    },
    gridColumn: {
      auto: 'auto',
      'span-1': 'span 1 / span 1',
      'span-2': 'span 2 / span 2',
      'span-3': 'span 3 / span 3',
      'span-4': 'span 4 / span 4',
      'span-5': 'span 5 / span 5',
      'span-6': 'span 6 / span 6',
      'span-7': 'span 7 / span 7',
      'span-8': 'span 8 / span 8',
      'span-9': 'span 9 / span 9',
      'span-10': 'span 10 / span 10',
      'span-11': 'span 11 / span 11',
      'span-12': 'span 12 / span 12'
    },
    gridColumnStart: {
      auto: 'auto',
      1: '1',
      2: '2',
      3: '3',
      4: '4',
      5: '5',
      6: '6',
      7: '7',
      8: '8',
      9: '9',
      10: '10',
      11: '11',
      12: '12',
      13: '13'
    },
    gridColumnEnd: {
      auto: 'auto',
      1: '1',
      2: '2',
      3: '3',
      4: '4',
      5: '5',
      6: '6',
      7: '7',
      8: '8',
      9: '9',
      10: '10',
      11: '11',
      12: '12',
      13: '13'
    },
    gridTemplateRows: {
      none: 'none',
      1: 'repeat(1, minmax(0, 1fr))',
      2: 'repeat(2, minmax(0, 1fr))',
      3: 'repeat(3, minmax(0, 1fr))',
      4: 'repeat(4, minmax(0, 1fr))',
      5: 'repeat(5, minmax(0, 1fr))',
      6: 'repeat(6, minmax(0, 1fr))'
    },
    gridRow: {
      auto: 'auto',
      'span-1': 'span 1 / span 1',
      'span-2': 'span 2 / span 2',
      'span-3': 'span 3 / span 3',
      'span-4': 'span 4 / span 4',
      'span-5': 'span 5 / span 5',
      'span-6': 'span 6 / span 6'
    },
    gridRowStart: {
      auto: 'auto',
      1: '1',
      2: '2',
      3: '3',
      4: '4',
      5: '5',
      6: '6',
      7: '7'
    },
    gridRowEnd: {
      auto: 'auto',
      1: '1',
      2: '2',
      3: '3',
      4: '4',
      5: '5',
      6: '6',
      7: '7'
    },
    height: theme => ({
      auto: 'auto',
      ...theme('spacing'),
      full: '100%',
      screen: '100vh'
    }),
    inset: {
      0: '0',
      1: '0.25rem',
      2: '0.5rem',
      3: '0.75rem',
      4: '1rem',
      5: '1.25rem',
      6: '1.5rem',
      8: '2rem',
      '-1': '-0.25rem',
      '-2': '-0.5rem',
      '-3': '-0.75rem',
      '-4': '-1rem',
      '-5': '-1.25rem',
      '-6': '-1.5rem',
      '-8': '-2rem',
      auto: 'auto'
    },
    letterSpacing: {
      tighter: '-0.05em',
      tight: '-0.025em',
      normal: '0',
      wide: '0.025em',
      wider: '0.05em',
      widest: '0.1em'
    },
    lineHeight: {
      none: '1',
      tight: '1.25',
      snug: '1.375',
      normal: '1.5',
      relaxed: '1.625',
      loose: '2'
    },
    listStyleType: {
      none: 'none',
      disc: 'disc',
      decimal: 'decimal'
    },
    margin: (theme, { negative }) => ({
      auto: 'auto',
      ...theme('spacing'),
      ...negative(theme('spacing')),
      ...negative({
        page: 'var(--padding-page)'
      })
    }),
    maxHeight: {
      full: '100%',
      screen: '100vh'
    },
    maxWidth: {
      xxs: '18rem',
      xs: '20rem',
      sm: '24rem',
      md: '28rem',
      lg: '32rem',
      xl: '36rem',
      '2xl': '42rem',
      '3xl': '48rem',
      '4xl': '56rem',
      '5xl': '64rem',
      '6xl': '72rem',
      full: '100%'
    },
    minHeight: {
      0: '0',
      full: '100%',
      screen: '100vh'
    },
    minWidth: {
      0: '0',
      full: '100%'
    },
    objectPosition: {
      bottom: 'bottom',
      center: 'center',
      left: 'left',
      'left-bottom': 'left bottom',
      'left-top': 'left top',
      right: 'right',
      'right-bottom': 'right bottom',
      'right-top': 'right top',
      top: 'top'
    },
    opacity: {
      0: '0',
      25: '0.25',
      50: '0.5',
      75: '0.75',
      100: '1'
    },
    order: {
      first: '-9999',
      last: '9999',
      none: '0',
      1: '1',
      2: '2',
      3: '3',
      4: '4',
      5: '5',
      6: '6',
      7: '7',
      8: '8',
      9: '9',
      10: '10',
      11: '11',
      12: '12'
    },
    padding: theme => theme('spacing'),
    placeholderColor: theme => theme('colors'),
    stroke: {
      current: 'currentColor'
    },
    textColor: theme => ({
      secondary: 'var(--text-secondary)',
      hint: 'var(--text-hint)',
      ...theme('colors'),
      'primary-contrast': {
        50: 'var(--color-primary-contrast-50)',
        100: 'var(--color-primary-contrast-100)',
        200: 'var(--color-primary-contrast-200)',
        300: 'var(--color-primary-contrast-300)',
        400: 'var(--color-primary-contrast-400)',
        500: 'var(--color-primary-contrast-500)',
        600: 'var(--color-primary-contrast-600)',
        700: 'var(--color-primary-contrast-700)',
        800: 'var(--color-primary-contrast-800)',
        900: 'var(--color-primary-contrast-900)'
      }
    }),
    width: theme => ({
      auto: 'auto',
      ...theme('spacing'),
      '1/2': '50%',
      '1/3': '33.333333%',
      '2/3': '66.666667%',
      '1/4': '25%',
      '2/4': '50%',
      '3/4': '75%',
      '1/5': '20%',
      '2/5': '40%',
      '3/5': '60%',
      '4/5': '80%',
      '1/6': '16.666667%',
      '2/6': '33.333333%',
      '3/6': '50%',
      '4/6': '66.666667%',
      '5/6': '83.333333%',
      '1/12': '8.333333%',
      '2/12': '16.666667%',
      '3/12': '25%',
      '4/12': '33.333333%',
      '5/12': '41.666667%',
      '6/12': '50%',
      '7/12': '58.333333%',
      '8/12': '66.666667%',
      '9/12': '75%',
      '10/12': '83.333333%',
      '11/12': '91.666667%',
      full: '100%',
      screen: '100vw'
    }),
    zIndex: {
      auto: 'auto',
      0: '0',
      10: '10',
      20: '20',
      30: '30',
      40: '40',
      50: '50'
    }
  },
  variants: {
    accessibility: ['responsive', 'focus'],
    alignContent: ['responsive'],
    alignItems: ['responsive'],
    alignSelf: ['responsive'],
    appearance: ['responsive'],
    backgroundAttachment: ['responsive'],
    backgroundColor: ['responsive', 'hover', 'focus'],
    backgroundPosition: ['responsive'],
    backgroundRepeat: ['responsive'],
    backgroundSize: ['responsive'],
    borderCollapse: ['responsive'],
    borderColor: ['responsive', 'hover', 'focus'],
    borderRadius: ['responsive'],
    borderStyle: ['responsive'],
    borderWidth: ['responsive', 'ltr', 'rtl'],
    boxShadow: ['responsive', 'hover', 'focus'],
    cursor: ['responsive'],
    display: ['responsive'],
    fill: ['responsive'],
    // Flex
    flex: ['responsive'],
    flexDirection: ['responsive'],
    flexGrow: ['responsive'],
    flexShrink: ['responsive'],
    flexWrap: ['responsive'],
    float: ['responsive'],
    fontFamily: ['responsive'],
    fontSize: ['responsive'],
    fontSmoothing: ['responsive'],
    fontStyle: ['responsive'],
    fontWeight: ['responsive', 'hover', 'focus'],
    // Grid
    gridAutoFlow: ['responsive'],
    gridTemplateColumns: ['responsive'],
    gridAutoColumns: ['responsive'],
    gridColumn: ['responsive'],
    gridColumnStart: ['responsive'],
    gridColumnEnd: ['responsive'],
    gridTemplateRows: ['responsive'],
    gridAutoRows: ['responsive'],
    gridRow: ['responsive'],
    gridRowStart: ['responsive'],
    gridRowEnd: ['responsive'],
    height: ['responsive'],
    inset: ['responsive'],
    justifyContent: ['responsive'],
    letterSpacing: ['responsive'],
    lineHeight: ['responsive'],
    listStylePosition: ['responsive'],
    listStyleType: ['responsive'],
    margin: ['responsive', 'ltr', 'rtl'],
    maxHeight: ['responsive'],
    maxWidth: ['responsive'],
    minHeight: ['responsive'],
    minWidth: ['responsive'],
    objectFit: ['responsive'],
    objectPosition: ['responsive'],
    opacity: ['responsive', 'hover', 'focus'],
    order: ['responsive'],
    outline: ['responsive', 'focus'],
    overflow: ['responsive'],
    padding: ['responsive', 'ltr', 'rtl'],
    placeholderColor: ['responsive', 'focus'],
    pointerEvents: ['responsive'],
    position: ['responsive'],
    resize: ['responsive'],
    stroke: ['responsive'],
    tableLayout: ['responsive'],
    textAlign: ['responsive'],
    textColor: ['responsive', 'hover', 'focus'],
    textDecoration: ['responsive', 'hover', 'focus'],
    textTransform: ['responsive'],
    userSelect: ['responsive'],
    verticalAlign: ['responsive'],
    visibility: ['responsive'],
    whitespace: ['responsive'],
    width: ['responsive'],
    wordBreak: ['responsive'],
    zIndex: ['responsive']
  },
  corePlugins: {},
  plugins: [
    function ({ addVariant, e }) {
      addVariant('ltr', ({ separator, modifySelectors }) => {
        modifySelectors(({ className }) => {
          return `[dir=ltr] .ltr${e(separator)}${className}`;
        });
      });

      addVariant('rtl', ({ separator, modifySelectors }) => {
        modifySelectors(({ className }) => {
          return `[dir=rtl] .rtl${e(separator)}${className}`;
        });
      });
    }
  ]
};
