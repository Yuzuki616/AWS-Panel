// Styles
import "../../../src/components/VTextField/VTextField.sass";
import "../../../src/components/VOtpInput/VOtpInput.sass"; // Extensions

import VInput from '../VInput';
import VTextField from '../VTextField/VTextField'; // Directives

import ripple from '../../directives/ripple'; // Utilities

import { convertToUnit, keyCodes } from '../../util/helpers';
import { breaking } from '../../util/console'; // Types

import mixins from '../../util/mixins';
const baseMixins = mixins(VInput);
/* @vue/component */

export default baseMixins.extend().extend({
  name: 'v-otp-input',
  directives: {
    ripple
  },
  inheritAttrs: false,
  props: {
    length: {
      type: [Number, String],
      default: 6
    },
    type: {
      type: String,
      default: 'text'
    },
    plain: Boolean
  },
  data: () => ({
    badInput: false,
    initialValue: null,
    isBooted: false,
    otp: []
  }),
  computed: {
    outlined() {
      return !this.plain;
    },

    classes() {
      return { ...VInput.options.computed.classes.call(this),
        ...VTextField.options.computed.classes.call(this),
        'v-otp-input--plain': this.plain
      };
    },

    isDirty() {
      return VInput.options.computed.isDirty.call(this) || this.badInput;
    }

  },
  watch: {
    isFocused: 'updateValue',

    value(val) {
      this.lazyValue = val;
      this.otp = (val == null ? void 0 : val.split('')) || [];
    }

  },

  created() {
    var _this$internalValue;

    /* istanbul ignore next */
    if (this.$attrs.hasOwnProperty('browser-autocomplete')) {
      breaking('browser-autocomplete', 'autocomplete', this);
    }

    this.otp = ((_this$internalValue = this.internalValue) == null ? void 0 : _this$internalValue.split('')) || [];
  },

  mounted() {
    requestAnimationFrame(() => this.isBooted = true);
  },

  methods: {
    /** @public */
    focus(e, otpIdx) {
      this.onFocus(e, otpIdx || 0);
    },

    genInputSlot(otpIdx) {
      return this.$createElement('div', this.setBackgroundColor(this.backgroundColor, {
        staticClass: 'v-input__slot',
        style: {
          height: convertToUnit(this.height)
        },
        on: {
          click: () => this.onClick(otpIdx),
          mousedown: e => this.onMouseDown(e, otpIdx),
          mouseup: e => this.onMouseUp(e, otpIdx)
        }
      }), [this.genDefaultSlot(otpIdx)]);
    },

    genControl(otpIdx) {
      return this.$createElement('div', {
        staticClass: 'v-input__control'
      }, [this.genInputSlot(otpIdx)]);
    },

    genDefaultSlot(otpIdx) {
      return [this.genFieldset(), this.genTextFieldSlot(otpIdx)];
    },

    genContent() {
      return Array.from({
        length: +this.length
      }, (_, i) => {
        return this.$createElement('div', this.setTextColor(this.validationState, {
          staticClass: 'v-input',
          class: this.classes
        }), [this.genControl(i)]);
      });
    },

    genFieldset() {
      return this.$createElement('fieldset', {
        attrs: {
          'aria-hidden': true
        }
      }, [this.genLegend()]);
    },

    genLegend() {
      const span = this.$createElement('span', {
        domProps: {
          innerHTML: '&#8203;'
        }
      });
      return this.$createElement('legend', {
        style: {
          width: '0px'
        }
      }, [span]);
    },

    genInput(otpIdx) {
      const listeners = Object.assign({}, this.listeners$);
      delete listeners.change; // Change should not be bound externally

      return this.$createElement('input', {
        style: {},
        domProps: {
          value: this.otp[otpIdx],
          min: this.type === 'number' ? 0 : null
        },
        attrs: { ...this.attrs$,
          disabled: this.isDisabled,
          readonly: this.isReadonly,
          type: this.type,
          id: `${this.computedId}--${otpIdx}`,
          class: `otp-field-box--${otpIdx}`,
          maxlength: 1
        },
        on: Object.assign(listeners, {
          blur: this.onBlur,
          input: e => this.onInput(e, otpIdx),
          focus: e => this.onFocus(e, otpIdx),
          paste: e => this.onPaste(e, otpIdx),
          keydown: this.onKeyDown,
          keyup: e => this.onKeyUp(e, otpIdx)
        }),
        ref: 'input',
        refInFor: true
      });
    },

    genTextFieldSlot(otpIdx) {
      return this.$createElement('div', {
        staticClass: 'v-text-field__slot'
      }, [this.genInput(otpIdx)]);
    },

    onBlur(e) {
      this.isFocused = false;
      e && this.$nextTick(() => this.$emit('blur', e));
    },

    onClick(otpIdx) {
      if (this.isFocused || this.isDisabled || !this.$refs.input[otpIdx]) return;
      this.onFocus(undefined, otpIdx);
    },

    onFocus(e, otpIdx) {
      e == null ? void 0 : e.preventDefault();
      e == null ? void 0 : e.stopPropagation();
      const elements = this.$refs.input;
      const ref = this.$refs.input && elements[otpIdx || 0];
      if (!ref) return;

      if (document.activeElement !== ref) {
        ref.focus();
        return ref.select();
      }

      if (!this.isFocused) {
        this.isFocused = true;
        ref.select();
        e && this.$emit('focus', e);
      }
    },

    onInput(e, otpIdx) {
      const target = e.target;
      const value = target.value;
      this.applyValue(otpIdx, target.value, () => {
        this.internalValue = this.otp.join('');
      });
      this.badInput = target.validity && target.validity.badInput;
      const nextIndex = otpIdx + 1;

      if (value) {
        if (nextIndex < +this.length) {
          this.changeFocus(nextIndex);
        } else {
          this.clearFocus(otpIdx);
          this.onCompleted();
        }
      }
    },

    clearFocus(index) {
      const input = this.$refs.input[index];
      input.blur();
    },

    onKeyDown(e) {
      if (e.keyCode === keyCodes.enter) {
        this.$emit('change', this.internalValue);
      }

      this.$emit('keydown', e);
    },

    onMouseDown(e, otpIdx) {
      // Prevent input from being blurred
      if (e.target !== this.$refs.input[otpIdx]) {
        e.preventDefault();
        e.stopPropagation();
      }

      VInput.options.methods.onMouseDown.call(this, e);
    },

    onMouseUp(e, otpIdx) {
      if (this.hasMouseDown) this.focus(e, otpIdx);
      VInput.options.methods.onMouseUp.call(this, e);
    },

    onPaste(event, index) {
      var _event$clipboardData;

      const maxCursor = +this.length - 1;
      const inputVal = event == null ? void 0 : (_event$clipboardData = event.clipboardData) == null ? void 0 : _event$clipboardData.getData('Text');
      const inputDataArray = (inputVal == null ? void 0 : inputVal.split('')) || [];
      event.preventDefault();
      const newOtp = [...this.otp];

      for (let i = 0; i < inputDataArray.length; i++) {
        const appIdx = index + i;
        if (appIdx > maxCursor) break;
        newOtp[appIdx] = inputDataArray[i].toString();
      }

      this.otp = newOtp;
      this.internalValue = this.otp.join('');
      const targetFocus = Math.min(index + inputDataArray.length, maxCursor);
      this.changeFocus(targetFocus);

      if (newOtp.length === +this.length) {
        this.onCompleted();
        this.clearFocus(targetFocus);
      }
    },

    applyValue(index, inputVal, next) {
      const newOtp = [...this.otp];
      newOtp[index] = inputVal;
      this.otp = newOtp;
      next();
    },

    changeFocus(index) {
      this.onFocus(undefined, index || 0);
    },

    updateValue(val) {
      // Sets validationState from validatable
      this.hasColor = val;

      if (val) {
        this.initialValue = this.lazyValue;
      } else if (this.initialValue !== this.lazyValue) {
        this.$emit('change', this.lazyValue);
      }
    },

    onKeyUp(event, index) {
      event.preventDefault();
      const eventKey = event.key;

      if (['Tab', 'Shift', 'Meta', 'Control', 'Alt'].includes(eventKey)) {
        return;
      }

      if (['Delete'].includes(eventKey)) {
        return;
      }

      if (eventKey === 'ArrowLeft' || eventKey === 'Backspace' && !this.otp[index]) {
        return index > 0 && this.changeFocus(index - 1);
      }

      if (eventKey === 'ArrowRight') {
        return index + 1 < +this.length && this.changeFocus(index + 1);
      }
    },

    onCompleted() {
      const rsp = this.otp.join('');

      if (rsp.length === +this.length) {
        this.$emit('finish', rsp);
      }
    }

  },

  render(h) {
    return h('div', {
      staticClass: 'v-otp-input',
      class: this.themeClasses
    }, this.genContent());
  }

});
//# sourceMappingURL=VOtpInput.js.map