"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = void 0;

var _vue = _interopRequireDefault(require("vue"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

var _default = _vue.default.extend({
  name: 'mouse',
  methods: {
    getDefaultMouseEventHandlers: function getDefaultMouseEventHandlers(suffix, getEvent) {
      var _this$getMouseEventHa;

      return this.getMouseEventHandlers((_this$getMouseEventHa = {}, _defineProperty(_this$getMouseEventHa, 'click' + suffix, {
        event: 'click'
      }), _defineProperty(_this$getMouseEventHa, 'contextmenu' + suffix, {
        event: 'contextmenu',
        prevent: true,
        result: false
      }), _defineProperty(_this$getMouseEventHa, 'mousedown' + suffix, {
        event: 'mousedown'
      }), _defineProperty(_this$getMouseEventHa, 'mousemove' + suffix, {
        event: 'mousemove'
      }), _defineProperty(_this$getMouseEventHa, 'mouseup' + suffix, {
        event: 'mouseup'
      }), _defineProperty(_this$getMouseEventHa, 'mouseenter' + suffix, {
        event: 'mouseenter'
      }), _defineProperty(_this$getMouseEventHa, 'mouseleave' + suffix, {
        event: 'mouseleave'
      }), _defineProperty(_this$getMouseEventHa, 'touchstart' + suffix, {
        event: 'touchstart'
      }), _defineProperty(_this$getMouseEventHa, 'touchmove' + suffix, {
        event: 'touchmove'
      }), _defineProperty(_this$getMouseEventHa, 'touchend' + suffix, {
        event: 'touchend'
      }), _this$getMouseEventHa), getEvent);
    },
    getMouseEventHandlers: function getMouseEventHandlers(events, getEvent) {
      var _this = this;

      var on = {};

      var _loop = function _loop(event) {
        var eventOptions = events[event];
        if (!_this.$listeners[event]) return "continue"; // TODO somehow pull in modifiers

        var prefix = eventOptions.passive ? '&' : (eventOptions.once ? '~' : '') + (eventOptions.capture ? '!' : '');
        var key = prefix + eventOptions.event;

        var handler = function handler(e) {
          var mouseEvent = e;

          if (eventOptions.button === undefined || mouseEvent.buttons > 0 && mouseEvent.button === eventOptions.button) {
            if (eventOptions.prevent) {
              e.preventDefault();
            }

            if (eventOptions.stop) {
              e.stopPropagation();
            } // Due to TouchEvent target always returns the element that is first placed
            // Even if touch point has since moved outside the interactive area of that element
            // Ref: https://developer.mozilla.org/en-US/docs/Web/API/Touch/target
            // This block of code aims to make sure touchEvent is always dispatched from the element that is being pointed at


            if (e && 'touches' in e) {
              var _e$currentTarget, _e$target;

              var classSeparator = ' ';
              var eventTargetClasses = (_e$currentTarget = e.currentTarget) == null ? void 0 : _e$currentTarget.className.split(classSeparator);
              var currentTargets = document.elementsFromPoint(e.changedTouches[0].clientX, e.changedTouches[0].clientY); // Get "the same kind" current hovering target by checking
              // If element has the same class of initial touch start element (which has touch event listener registered)

              var currentTarget = currentTargets.find(function (t) {
                return t.className.split(classSeparator).some(function (c) {
                  return eventTargetClasses.includes(c);
                });
              });

              if (currentTarget && !((_e$target = e.target) != null && _e$target.isSameNode(currentTarget))) {
                currentTarget.dispatchEvent(new TouchEvent(e.type, {
                  changedTouches: e.changedTouches,
                  targetTouches: e.targetTouches,
                  touches: e.touches
                }));
                return;
              }
            }

            _this.$emit(event, getEvent(e), e);
          }

          return eventOptions.result;
        };

        if (key in on) {
          /* istanbul ignore next */
          if (Array.isArray(on[key])) {
            on[key].push(handler);
          } else {
            on[key] = [on[key], handler];
          }
        } else {
          on[key] = handler;
        }
      };

      for (var event in events) {
        var _ret = _loop(event);

        if (_ret === "continue") continue;
      }

      return on;
    }
  }
});

exports.default = _default;
//# sourceMappingURL=mouse.js.map